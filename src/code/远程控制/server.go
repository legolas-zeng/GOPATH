package main

import (
    "fmt"
    "github.com/whuwzp/RemoteControl/RemoteControl-2.0/message"
    "net"
    "sync"
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
    slaveConnPool []net.Conn
    masterConn    net.Conn
    mu            sync.Mutex
)

func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:5000")
    if err != nil {
        panic(err)
    }
    for {
        conn, err := listen.Accept()
        if err != nil {
            panic(err)
        }
        fmt.Println("========================\na new conn: ", conn.RemoteAddr().String())
        AddslaveConnPool(conn)

        go handleConn(conn)
    }
    defer listen.Close()

}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        buf := make([]byte, bufsize)
        RecvMsg := message.RecvMsg(c, buf)

        //delete the master conn from pool
        if RecvMsg.From == "master" {
            masterConn = c
            DelslaveConnPool(c)
        }

        dispatch(RecvMsg)

    }
}

func AddslaveConnPool(c net.Conn) {
    mu.Lock()
    defer mu.Unlock()
    slaveConnPool = append(slaveConnPool, c)
    fmt.Println("conn pool :", slaveConnPool)
}

func DelslaveConnPool(c net.Conn) {
    mu.Lock()
    defer mu.Unlock()
    //fmt.Println("pool before delete", slaveConnPool)
    var pool []net.Conn
    for _, v := range slaveConnPool {
        if v == c {

        } else {
            pool = append(pool, v)
        }
    }
    slaveConnPool = pool
    //fmt.Println("pool after delete", slaveConnPool)
}

func dispatch(m message.Message) {
    if m.To == MASTER {
        slave2master(m)
    } else if m.To == SERVER {
        master2server(m)
    } else {
        master2slave(m)
    }
}

func slave2master(m message.Message) {
    message.SendMsg(masterConn, m)
}

func master2slave(m message.Message) {
    for _, c := range slaveConnPool {
        if c.RemoteAddr().String() == m.To {
            if m.Type == TYPEDATA{
                message.SendMsg(c, m)
                message.TransFile(c, masterConn)
            } else {
                message.SendMsg(c, m)
            }
            break
        }

    }
}
func master2server(m message.Message) {
    var msg message.Message
    if m.Cmd == "listslave" {
        msg = message.Message{
            SERVER,
            MASTER,
            TYPEDATA,
            "",
            string(listslave()),
        }
    }
    message.SendMsg(masterConn, msg)
}

func listslave() string {
    mu.Lock()
    defer mu.Unlock()
    var list string
    for _, c := range slaveConnPool {
        list += (c.RemoteAddr().String() + " ")
    }
    fmt.Println("list:", list)
    return list
}