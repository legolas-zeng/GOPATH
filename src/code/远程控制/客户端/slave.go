package 客户端

import (
    "fmt"
    "net"

    "log"
    "os/exec"

    "bufio"
    "io"
    "os"

    "github.com/whuwzp/RemoteControl/RemoteControl-2.0/message"
)

const (
   bufsize  = 4096 * 100
   MASTER   = "master"
   SLAVE    = "slave"
   SERVER   = "server"
   TYPECMD  = "cmd"
   TYPECODE = "code"
   TYPEDATA = "data"
)

var (
    DATA = make([]string, 4096)
    conn net.Conn
)

func main() {
    var err error
    conn, err = net.Dial("tcp", "127.0.0.1:5000")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    buf := make([]byte, 4096)
    for {
        RecvMsg := message.RecvMsg(conn, buf)
        dispatchs(RecvMsg)
    }

}

func dispatchs(m message.Message) {
    if m.Type == TYPECMD {
        msg := execCommand(m.Cmd, m.Args)
        RespMsg := message.Message{
            From: SLAVE,
            To:   MASTER,
            Type: TYPECMD,
            Cmd:  "",
            Args: msg,
        }
        message.SendMsg(conn, RespMsg)
    } else if m.Type == TYPECODE {
        writeCommand(m.Args)
    } else {
        //file
        message.SendFile(conn, m.Args)
    }
}

//for example: python test.py

func execCommand(Cmd string, arg string) string {
    cmd := exec.Command(Cmd, arg)
    fmt.Println("going to exe command...")
    //显示运行的命令
    fmt.Println(cmd.Args)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println(err)
    }
    cmd.Start()
    reader := bufio.NewReader(stdout)
    msg := ""
    //实时循环读取输出流中的一行内容
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
            break
        }
        msg += line
    }
    fmt.Println(msg)
    cmd.Wait()
    return msg
}

func writeCommand(code string) {
    fmt.Println("going to write code...")
    f, err := os.Create("python.py")
    if err != nil {
        log.Println(err)
    }
    defer f.Close()

    f.WriteString(code)
    fmt.Println("writing code:", code)
    fmt.Println("writen!")
}

func fileCommand(filename string){

}