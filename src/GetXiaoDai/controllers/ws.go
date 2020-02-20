package controllers

import (
	"log"
	"time"
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"io"
	"github.com/hpcloud/tail"
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
	tailTask()
}

func tailTask() {
	fmt.Println("启动websocket！！！")
	go handleexcle()
	go tailMsg()

}

func tailMsg(){
	filename := "C:\\Users\\Administrator\\Desktop\\test.log"

	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		// Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		broadcast <- msg.Text
	}
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
	t1 := time.Now()
	for _, row := range sheet.Rows[1:] {
		number := row.Cells[0]
		filename := row.Cells[7]
		fullname := fmt.Sprintf("C:\\Users\\Administrator\\Desktop\\%s-%s", number, filename)
		srcpath := fmt.Sprintf("%s",row.Cells[8])

		//fmt.Println(srcpath,fullname)
		time.Sleep(2*time.Second)
		_ , err := copy(srcpath,fullname)
		if err == nil {
			log2fileAndStdout(fmt.Sprintf("success-----%s的%s：%s拷贝完成",row.Cells[2],row.Cells[5],row.Cells[7]))
		}else {
			log2fileAndStdout(fmt.Sprintf("fail-----%s的%s：%s拷贝完成",row.Cells[2],row.Cells[5],row.Cells[7]))
		}
	}
	log2fileAndStdout(fmt.Sprintf("------全部完成！------"))
	elapsed := time.Since(t1)
	log2fileAndStdout(fmt.Sprintf("------共计%d行,总共用时%s！------",len(sheet.Rows),elapsed))

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


func copy(src string, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

