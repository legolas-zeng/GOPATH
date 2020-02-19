package controllers

import (
	"log"

	"GetXiaoDai/models"
	"time"

	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *WebSocketController) Get() {

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true

	timeTask()

}

func timeTask() {

	timeStr := "0/3 * * * * *" //每隔3秒执行
	//toolbox 定时任务
	t1 := toolbox.NewTask("timeTask", timeStr, func() error {

		//todo do what you want
		msg := models.SocketMessage{SocketMessage: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg

		return nil
	})

	toolbox.AddTask("tk1", t1)
	toolbox.StartTask()
}
