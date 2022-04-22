package socket

import "github.com/gorilla/websocket"

// SocketEventStruct struct of socket events
type ISocketEvent struct {
	Kind    string      `json:"kind"`
	Payload interface{} `json:"payload"`
}

type IClient struct {
	hub        *IHub
	connection *websocket.Conn
	userId     string

	send chan ISocketEvent
}

type IJoinPayload struct {
	UserID string `json:"user_id"`
}

type ILeftPayload struct {
	UserID string `json:"user_id"`
}

// hub                 *IHub
// webSocketConnection *websocket.Conn
// send                chan SocketEventStruct
// username            string
// userID              string
