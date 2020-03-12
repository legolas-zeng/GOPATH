package main

import (
"fmt"
"io"
"net"
)

func main() {

    conn, err := net.Dial("tcp", "127.0.0.1:8082")
    if err != nil {
        fmt.Println("连接错误", err.Error())
        return
    }
    defer conn.Close()
    msg := "GET / HTTP/1.1\r\n"
    msg += "Host:127.0.0.1:8082\r\n"
    msg += "Connection:keep-alive\r\n"
    //msg += "User-Agent:Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36\r\n"
    msg += "\r\n\r\n"
    fmt.Println(msg)
    //io.WriteString(os.Stdout, msg)
    n, err := io.WriteString(conn, msg)
    if err != nil {
        fmt.Println("写入字符串失败, ", err)
        return
    }
    fmt.Println("发送数据:", n)
    buf := make([]byte, 4096)
    for {
        count, err := conn.Read(buf)
        fmt.Println("count:", count, "err:", err)
        if err != nil {
            break
        }
        fmt.Println(string(buf[0:count]))
    }
}

