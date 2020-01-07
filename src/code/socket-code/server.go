package main

import (
    "net"
    "log"
    "fmt"
)

func Listening() {
    tcpListen, err := net.Listen("tcp", ":8565")

    if err != nil {
        panic(err)
    }

    for {
        conn, err := tcpListen.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go connHandle(conn)
    }
}

func  connHandle(conn net.Conn) {
    defer conn.Close()
    readBuff := make([]byte, 14)
    for {
        n, err := conn.Read(readBuff)
        if err != nil {
            return
        }
        fmt.Println(readBuff[:n])

    }
}


func main(){
    Listening()
}
