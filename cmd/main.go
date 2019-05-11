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
	fmt.Println(c, err)
}
