package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type IClient struct {
	ID     string
	Conn   *websocket.Conn
	Pool   *Pool
	UserId string
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func UserLeft(c *IClient) {
	c.Pool.Left <- c
	c.Conn.Close()
}

func (c *IClient) Read() {
	defer func() {
		UserLeft(c)
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// eventDataString := string(p[:])
		// fmt.Printf("eventDataString = %v", eventDataString)

		var eventData ISocketEvent
		if err := json.Unmarshal(p, &eventData); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nEventData = %v", eventData)

		// switch eventData.Kind {

		// case KindTyping:
		// 	data := IUserTypeData{Kind: KindTyping, OuserId: c.UserId}
		// 	c.Pool.Broadcast <- data
		// 	break

		// }
	}
}
