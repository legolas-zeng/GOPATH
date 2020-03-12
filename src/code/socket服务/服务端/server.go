package main

//TODO 不支持多线程监听client，client无法持续监听server

import (
    "log"
    "net"
    "os"
    "os/signal"
)

func main() {
    stop_chan := make(chan os.Signal) // 接收系统中断信号
    signal.Notify(stop_chan, os.Interrupt)
    netListen, err := net.Listen("tcp", "0.0.0.0:2048")
    if err != nil {
        log.Println(err.Error())
        return
    }
    defer netListen.Close()

    log.Println("Waiting for clients")
    for {
        conn, err := netListen.Accept()
        if err != nil {
            log.Println(err.Error())
            continue
        }
        log.Println(conn.RemoteAddr().String(), " tcp 连接成功！")
        handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    buffer := make([]byte, 2048)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            log.Println(conn.RemoteAddr().String(), " read error: ", err.Error())
            return
        }
        log.Println(conn.RemoteAddr().String(), "接收到数据:", string(buffer[:n]))
    }
}