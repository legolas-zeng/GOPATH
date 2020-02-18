package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"GetXiaoDai/models"
	"time"
	//"github.com/astaxie/beego/toolbox"
)

type MyWebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *MyWebSocketController) Get() {

	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	//  defer ws.Close()

	clients[ws] = true

	//不断的广播发送到页面上
	for {
		//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
		time.Sleep(time.Second * 3)
		msg := models.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
	}
}