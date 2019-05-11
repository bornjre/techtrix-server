package app

import "github.com/gorilla/websocket"

type subscribers struct {
	//websockets
	AllConns []*websocket.Conn

	BroadcastChan chan string
}

func NewSubscriberService() *subscribers {
	return &subscribers{
		AllConns:      make([]*websocket.Conn, 10),
		BroadcastChan: make(chan string),
	}
}

func (b *subscribers) RunService() {
	for {
		select {
		case msg := <-b.BroadcastChan:
			b.Broadcast(msg)
		}
	}
}

func (b *subscribers) Broadcast(message string) {
	for _, conn := range b.AllConns {
		conn.WriteMessage(websocket.TextMessage, []byte(message))

	}
}
