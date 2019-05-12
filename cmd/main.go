package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Starting validator ..")
	urlstr := "ws://localhost:8080/subscribe"
	fmt.Println("Connecting to ..", urlstr)
	c, _, err := websocket.DefaultDialer.Dial(urlstr, nil)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(c, err)
	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			println(err)
			continue
		}
		println("HASH RECIVED")
		println(string(b))
	}
}
