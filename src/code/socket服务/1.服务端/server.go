package main

//TODO client无法持续监听

import (
    "fmt"
    "net"
)

func creatServer(){
    //创建服务器地址
    addr,_:=net.ResolveTCPAddr("tcp4","localhost:8899")
    //创建监听器
    lis,_:=net.ListenTCP("tcp4",addr)
    fmt.Println("服务器已启动")
    //通过监听器获取客户端传递过来的数据
    for{
        conn, _ := lis.Accept()
        go func() {
            //转换数据
            b:=make([]byte,1024)
            n,_:=conn.Read(b)
            fmt.Println("获取到的数据：",string(b[:n]))
            //发送消息
            conn.Write(append([]byte("server:"),b[:n]...))
            //关闭连接
            conn.Close()
        }()
    }

}
func main() {
    creatServer()
}