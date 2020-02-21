package controllers

import (
	"log"

	"github.com/gorilla/websocket"
	"fmt"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	//broadcast = make(chan models.SocketMessage)
	broadcast = make(chan string)
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		//fmt.Println("已连接的客户端数量：", len(clients))
		fmt.Println("-------",msg)
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
