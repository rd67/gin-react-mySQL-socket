package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rd67/gin-react-mySQL-socket/helpers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return conn, nil
}

func ServeWs(pool *Pool, ctx *gin.Context) {
	conn, err := upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+v\n", err)
		return
	}

	userId := helpers.GetAuthUserId(ctx)

	client := &IClient{
		Conn:   conn,
		Pool:   pool,
		UserId: userId,
	}

	pool.Join <- client
	client.Read()
}
