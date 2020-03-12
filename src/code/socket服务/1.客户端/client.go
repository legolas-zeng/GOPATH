package main

import (
    "fmt"
    "net"
)

func creatClient(){
    //创建服务器端地址
    addr,_:=net.ResolveTCPAddr("tcp4","localhost:8899")
    //创建连接
    conn, _ := net.DialTCP("tcp4", nil, addr)
    //发生数据
    conn.Write([]byte("客户端发送的数据"))
    b:=make([]byte,1024)
    count,_:=conn.Read(b)
    fmt.Println("服务器发送回来的消息为：",string(b[:count]))
    //关闭连接
    conn.Close()
}
func main()  {
    creatClient()
}
