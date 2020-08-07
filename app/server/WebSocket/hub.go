package WebSocket

import "log"

type Message struct {
	Id   string
	Data []byte
}

type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	Broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		Broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
			log.Println("register success :" + client.id)
			log.Println(h)
			//h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
				close(client.send)
			}
		case message := <-h.Broadcast:
			if client, ok := h.clients[message.Id]; ok {
				select {
				case client.send <- message.Data:
					log.Println("push success")
				default:
					close(client.send)
					delete(h.clients, message.Id)
				}
			}
		}
	}
}
