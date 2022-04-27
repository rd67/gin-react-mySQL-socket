package websocket

import "fmt"

type Pool struct {
	Join      chan *IClient
	Left      chan *IClient
	Clients   map[*IClient]bool
	Broadcast chan interface{} //	Sends to all clients
	Send      chan interface{} //	Sendts to single client
}

func NewPool() *Pool {
	return &Pool{
		Join:      make(chan *IClient),
		Left:      make(chan *IClient),
		Clients:   make(map[*IClient]bool),
		Broadcast: make(chan interface{}),
		Send:      make(chan interface{}),
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
					//	In Case of error removing the socket client
					UserLeft(client)
					fmt.Printf("Error while broadcasting to User:%s, %v", client.UserId, err)
					return
				}
			}
			break

		case message := <-pool.Send:
			for client := range pool.Clients {

				err := client.Conn.WriteJSON(message)
				if err != nil {
					UserLeft(client)
					fmt.Printf("\nError while emiting to User:%s, %v", client.UserId, err)
					return
				}

				fmt.Printf("\nEmited to User:%s, data:%v", client.UserId, message)

			}

			break

		}
	}
}
