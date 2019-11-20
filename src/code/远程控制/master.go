package main

import (
    "fmt"
    "net"
    //"regexp"
    //"os"
    //"encoding/json"
    "github.com/whuwzp/RemoteControl/RemoteControl-2.0/message"

)

const (
  MASTER   = "master"
  SLAVE    = "slave"
  SERVER   = "server"
  TYPECMD  = "cmd"
  TYPECODE = "code"
  TYPEDATA = "data"
)

var VoidCmd []string

func init() {
    VoidCmd = []string{"python", "calc"}
}

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:5000")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    buf := make([]byte, 4096)
    flag := false
    for {
        flag = false
        m := GetCmd()
        fmt.Println("starting to send msg...")
        for _, v := range VoidCmd {
            if m.Cmd == v {
                //sending
                message.SendMsg(conn, m)
                flag = true
                break
            }
        }
        if !flag {
            if m.Type == TYPEDATA{
                message.SendMsg(conn, m)
                fmt.Println("starting to recv file...")
                message.RecvFile(conn, m.Args)
            }else {
                message.SendMsg(conn, m)
                fmt.Println("starting to recv msg...")
                RecvMsg := message.RecvMsg(conn, buf)
                fmt.Println(RecvMsg.Args)
            }
        }
    }
}

func GetCmd() message.Message {
    var from, to, cmdtype, cmd, args string
    //from
    from = MASTER

    //to
    fmt.Println("please select your slave (or 0: server): ")
    fmt.Scanln(&to)
    if to == "0" {
        to = SERVER
    }

    //type
    fmt.Println("please select your command type: \n" +
        "0: execCmd\n" +
        "1: writeCode\n" +
        "2: Data")
    fmt.Scanln(&cmdtype)
    if cmdtype == "0" {
        cmdtype = TYPECMD
    } else if cmdtype == "1" {
        cmdtype = TYPECODE
    } else {
        cmdtype = TYPEDATA
    }

    //cmd
    fmt.Println("please input your command: (0: listslave)")
    fmt.Scanln(&cmd)
    if cmd == "0" {
        cmd = "listslave"
    }

    //args
    fmt.Println("please input your command args: ")
    fmt.Scanln(&args)

    return message.Message{from, to, cmdtype, cmd, args}
}