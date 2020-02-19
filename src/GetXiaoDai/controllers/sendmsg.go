package controllers

import (
	"fmt"
	"GetXiaoDai/models"
	"log"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.SocketMessage)
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Println("已连接的客户端数量：", len(clients))
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
