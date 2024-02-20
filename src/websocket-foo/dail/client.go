package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func dialFoo(urlStr string) {
	dialURL, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	// dialURL := url.URL{Scheme: "ws", Host: "192.168.2.203:6666", Path: "/dial"}
	log.Println("connecting to", dialURL.String())

	websocket.DefaultDialer.HandshakeTimeout = time.Minute
	c, _, err := websocket.DefaultDialer.Dial(dialURL.String(), nil)
	if err != nil {
		log.Fatalln("dial failed,", err)
	}
	defer c.Close()

	select {}

	// done := make(chan struct{})

	// go func() {
	// 	defer close(done)
	// 	for {
	// 		_, message, err := c.ReadMessage()
	// 		if err != nil {
	// 			log.Println("read:", err)
	// 			return
	// 		}
	// 		log.Printf("recv: %s", message)
	// 	}
	// }()

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case t := <-ticker.C:
	// 		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	// 		if err != nil {
	// 			log.Println("write:", err)
	// 			return
	// 		}
	// 	case <-interrupt:
	// 		log.Println("interrupt")

	// 		// Cleanly close the connection by sending a close message and then
	// 		// waiting (with timeout) for the server to close the connection.
	// 		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// 		if err != nil {
	// 			log.Println("write close:", err)
	// 			return
	// 		}
	// 		select {
	// 		case <-done:
	// 		case <-time.After(time.Second):
	// 		}
	// 		return
	// 	}
	// }
}
