package app

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type subscribers struct {
	//websockets
	AllConns []*websocket.Conn

	BroadcastChan chan string
}

func NewSubscriberService() *subscribers {
	return &subscribers{
		AllConns:      make([]*websocket.Conn, 10),
		BroadcastChan: make(chan string, 20),
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

func (b *subscribers) subscribe(w http.ResponseWriter, r *http.Request) {
	printl("websocket!!")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	b.AllConns = append(b.AllConns, conn)

}
