package utils

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func ConnectSocket(c *gin.Context) {
	//https://tutorialedge.net/projects/chat-system-in-go-and-react/part-2-simple-communication/

	//https://webdevelop.pro/blog/guide-creating-websocket-client-golang-using-mutex-and-channel
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error during connection upgradation:", err)
		return
	}

	defer conn.Close()

	// helpful log statement to show connections
	log.Println("Client Connected")

	reader(conn)

}
