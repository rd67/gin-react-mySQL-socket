package socket

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rd67/gin-react-mySQL-socket/helpers"
)

// var upgrader = websocket.Upgrader{}

func ConnectSocket(c *gin.Context) {
	var upgrader = websocket.Upgrader{}

	var hub = NewHub()
	go hub.Run()

	userId := helpers.GetAuthUserId(c)

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!", userId)

	connection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading socket connection", err)
		return
	}

	ConnectSocketNewUser(hub, connection, userId)
}

func ConnectSocketNewUser(hub *IHub, connection *websocket.Conn, userId string) {
	client := &IClient{
		hub:        hub,
		connection: connection,
		userId:     userId,

		send: make(chan ISocketEvent),
	}

	go client.readPump()

	client.hub.connect <- client
}
