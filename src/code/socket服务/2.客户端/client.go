package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func handleErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    }
}

func main() {
    conn, err := net.Dial("tcp", ":20000")
    handleErr(err)
    defer conn.Close()
    reader := bufio.NewReader(os.Stdin)
    for {
        data, err := reader.ReadString('\n')
        handleErr(err)

        data = strings.TrimSpace(data)  // 简单来说以TrimSpace作为去除首尾的空白字符,在这里等效于Trim( data , "\r\n") ! 换行是CRLF模式
        fmt.Println("len" , len(data))
        if strings.ToUpper(data) == "END" {
            fmt.Println("END")
            conn.Write([]byte("end"))
            break
        }
        if len(data) == 0{
            continue
        }
        _, err = conn.Write([]byte(data))
        handleErr(err)
        var buf = [1024]byte{}
        n ,err := conn.Read(buf[:])
        handleErr(err)
        fmt.Println(string(buf[:n]) , n)
    }
}