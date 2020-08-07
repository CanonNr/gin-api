package WebSocket

import "log"

func Send(message Message) {
	h := HubServers
	log.Println(h.clients)

	if client, ok := h.clients[message.Id]; ok {
		select {
		case client.send <- message.Data:
			log.Println("push success")
		default:
			close(client.send)
			delete(h.clients, message.Id)
		}
	} else {
		log.Println("Client Error")
	}
}
