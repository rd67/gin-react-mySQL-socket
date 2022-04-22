package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	// "time"

	"github.com/gorilla/websocket"
)

// const (
// 	writeWait      = 10 * time.Second
// 	pongWait       = 60 * time.Second
// 	pingPeriod     = (pongWait * 9) / 10
// 	maxMessageSize = 512
// )

func HandleUserConnectEvent(hub *IHub, client *IClient) {
	hub.clients[client] = true

	fmt.Println("Socket Connected with userId: ", client.userId)
	//	This is To Fire Join Event Accross All Connected Connections/
	// TODO: Impletent broadcasing to only user contacts which are connected
	handleSocketPayloadEvents(client, ISocketEvent{
		Kind:    "Join",
		Payload: client.userId,
	})
}

func HandleUserDisconnectEvent(hub *IHub, client *IClient) {
	_, ok := hub.clients[client]
	if ok {
		delete(hub.clients, client)
		close(client.send)

		handleSocketPayloadEvents(client, ISocketEvent{
			Kind:    "Left",
			Payload: client.userId,
		})
	}

	fmt.Println("Socket DisConnect with userId: ", client.userId)
}

func handleSocketPayloadEvents(client *IClient, socketEvent ISocketEvent) {

	switch socketEvent.Kind {

	case "Join":
		log.Printf("Join Event Trigerred")
		BroadcastSocketEventToAllClient(client.hub, ISocketEvent{
			Kind: socketEvent.Kind,
			Payload: IJoinPayload{
				UserID: client.userId,
			},
		})
	case "Left":
		log.Printf("Left Event Trigerred")
		BroadcastSocketEventToAllClient(client.hub, ISocketEvent{
			Kind: socketEvent.Kind,
			Payload: ILeftPayload{
				UserID: client.userId,
			},
		})
	}

}

// BroadcastSocketEventToAllClient will emit the socket events to all socket users
func BroadcastSocketEventToAllClient(hub *IHub, payload ISocketEvent) {
	for client := range hub.clients {
		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, client)
		}
	}
}

// It will emit the socket event to specific socket user
func EmitToSpecificClient(hub *IHub, payload ISocketEvent, userId string) {
	for client := range hub.clients {
		if client.userId == userId {
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}

// func EmitToSpecificClient(hub *IHub, payload SocketEventStruct, userID string) {
// 	for client := range hub.clients {
// 		if client.userID == userID {
// 			select {
// 			case client.send <- payload:
// 			default:
// 				close(client.send)
// 				delete(hub.clients, client)
// 			}
// 		}
// 	}
// }

func unRegisterAndCloseConnection(c *IClient) {
	c.hub.disconnect <- c
	c.connection.Close()
}

// func setSocketPayloadReadConfig(c *IClient) {
// 	c.connection.SetReadLimit(maxMessageSize)
// 	c.connection.SetReadDeadline(time.Now().Add(pongWait))
// 	c.connection.SetPongHandler(func(string) error { c.connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
// }

func (c *IClient) readPump() {
	var socketEventPayload ISocketEvent

	defer unRegisterAndCloseConnection(c)

	// setSocketPayloadReadConfig(c)

	for {
		_, payload, err := c.connection.ReadMessage()

		decoder := json.NewDecoder(bytes.NewReader(payload))
		decoderErr := decoder.Decode(&socketEventPayload)

		if decoderErr != nil {
			log.Printf("error: %v", decoderErr)
			break
		}

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error ===: %v", err)
			}
			break
		}

		handleSocketPayloadEvents(c, socketEventPayload)
	}
}
