package socket

// Hub maintains the set of active clients and broadcasts messages to the clients
type IHub struct {
	clients map[*IClient]bool

	connect    chan *IClient
	disconnect chan *IClient
}

//	It will give a instance of a Hub
func NewHub() *IHub {
	return &IHub{
		clients: make(map[*IClient]bool),

		connect:    make(chan *IClient),
		disconnect: make(chan *IClient),
	}
}

//	It will execute Go Routines to check incoming Socket Events
func (hub *IHub) Run() {
	for {
		select {
		case client := <-hub.connect:
			HandleUserConnectEvent(hub, client)

		case client := <-hub.disconnect:
			HandleUserDisconnectEvent(hub, client)
		}
	}
}
