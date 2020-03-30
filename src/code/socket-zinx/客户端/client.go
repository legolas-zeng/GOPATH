package main

import (
    "fmt"
    "net"
    "time"
    "errors"
    "github.com/aceld/zinx/znet"
    "io"
)

/*
    模拟客户端
 */
func main() {

    //fmt.Println("Client Test ... start")
    ////3秒之后发起测试请求，给服务端开启服务的机会
    //time.Sleep(3 * time.Second)
    service:="192.168.10.3:8999"
    tcpAddr,_:=net.ResolveTCPAddr("tcp4",service)
    //conn, err := net.Dial("tcp", "192.168.10.3:8999")
    //if err != nil {
    //    fmt.Println("客户端启动失败!")
    //    return
    //}

    for {
        //发封包message消息
        conn,err:=net.DialTCP("tcp",nil,tcpAddr)
        if err != nil {
            fmt.Printf("未连接服务端:%s \n", err)
        } else {
            fmt.Printf("连接服务端:%s \n", conn.RemoteAddr().String())
            defer conn.Close()
            doWork(conn)
        }

        time.Sleep(1 * time.Second)
    }
}

//解决断线重连问题
func doWork(conn net.Conn) error {
    ch := make(chan int, 100)
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for {
        select {
        case stat := <-ch:
            if stat == 2 {
                return errors.New("None Msg")
            }
        case <-ticker.C:
            ch <- 1
            go ClientMsgHandler(conn, ch)
        case <-time.After(time.Second * 10):
            defer conn.Close()
            fmt.Println("超时断线....")
        }
    }
    return nil
}

func ClientMsgHandler(conn net.Conn, ch chan int) {
    <-ch
    //获取当前时间
    msg := "❤❤❤❤❤❤"
    go SendMsg(conn, msg)
    go ReadMsg(conn)
}

func SendMsg(conn net.Conn, msg string) {
    dp := znet.NewDataPack()
    msgs, _ := dp.Pack(znet.NewMsgPackage(0, []byte(msg)))
    conn.Write(msgs)
}

//接收服务端发来的消息
func ReadMsg(conn net.Conn) {
    dp := znet.NewDataPack()
    headData := make([]byte, dp.GetHeadLen())
    io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
    msgHead, err := dp.Unpack(headData)
    if err != nil {
        fmt.Println("server unpack err:", err)
        return
    }

    if msgHead.GetDataLen() > 0 {
        //msg 是有data数据的，需要再次读取data数据
        msg := msgHead.(*znet.Message)
        msg.Data = make([]byte, msg.GetDataLen())

        //根据dataLen从io中读取字节流
        _, err := io.ReadFull(conn, msg.Data)
        if err != nil {
            fmt.Println("server unpack data err:", err)
            return
        }

        fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))

    }
}
