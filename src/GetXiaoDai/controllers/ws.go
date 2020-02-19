package controllers

import (
	"log"

	"GetXiaoDai/models"
	"time"

	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/gorilla/websocket"
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"io"
)

type WebSocketController struct {
	beego.Controller
}

var (
	upgrader = websocket.Upgrader{}
	inFile = "C:\\Users\\Administrator\\Desktop\\20200114.xlsx"
)

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
	handleexcle()
}

func handleexcle(){
	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sheet := xlFile.Sheets[0]
	fmt.Println("工作表名: ", sheet.Name)
	for _, row := range sheet.Rows[1:] {
		number := row.Cells[0]
		filename := row.Cells[7]
		//name := row.Cells[2]
		//path := row.Cells[8]
		fullname := fmt.Sprintf("%s-%s", number, filename)
		fmt.Println(fullname)
		log2fileAndStdout(fullname)
		//for _, cell := range row.Cells {
		//    fmt.Println(cell)
		//}
		fmt.Print("\n")
	}
	fmt.Println("\n读取成功")
	//func_log2file()
	//func_log2fileAndStdout()

}

func log2fileAndStdout(msg string) {
	//创建日志文件
	f, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	// 定义多个写入器
	writers := []io.Writer{f, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
	// 使用新的log对象，写入日志内容
	logger.Println(msg)
}
