package pkg

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

// "github.com/gin-gonic/gin"

func SocketConnect() *socketio.Server {
	socketServer := socketio.NewServer(nil)

	socketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!connected:", s.ID())
		return nil
	})

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatal("Socketio listen error: ", err)
		}
	}()
	// defer socketServer.Close()

	return socketServer
}

// func WrapH(h http.Handler) HandlerFunc {
//     return func(c *Context) {
//         h.ServeHTTP(c.Writer, c.Request)
//     }
// }
