package websocket

import "fmt"

type Pool struct {
	Join      chan *IClient
	Left      chan *IClient
	Clients   map[*IClient]bool
	Broadcast chan interface{}
}

func NewPool() *Pool {
	return &Pool{
		Join:      make(chan *IClient),
		Left:      make(chan *IClient),
		Clients:   make(map[*IClient]bool),
		Broadcast: make(chan interface{}),
	}
}

func (pool *Pool) Start() {
	for {
		select {

		case client := <-pool.Join:
			pool.Clients[client] = true

			data := IUserJoinData{
				Kind:   KindJoin,
				UserId: client.UserId,
			}

			fmt.Println("Join by User Id: ", client.UserId)

			for client := range pool.Clients {
				client.Conn.WriteJSON(data)
			}
			break

		case client := <-pool.Left:
			delete(pool.Clients, client)

			data := IUserLeaveData{
				Kind:   KindLeft,
				UserId: client.UserId,
			}

			fmt.Println("Left by User Id: ", client.UserId)

			for client := range pool.Clients {
				client.Conn.WriteJSON(data)
			}
			break

		case message := <-pool.Broadcast:
			fmt.Println("Broadcasting message to all clients in Pool", message)
			for client := range pool.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			break

		}
	}
}
