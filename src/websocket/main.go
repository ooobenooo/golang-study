package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		// &reply 是一个指针地址
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't reply")
			break
		}

		fmt.Println("Receive message from client: ", reply)
		msg := "Receive: " + reply

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

// go get golang.org/x/net/websocket
// 下载websocket包
func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":38080", nil); err != nil {
		log.Fatal("Listen port 38080 error", err)
	}
}
