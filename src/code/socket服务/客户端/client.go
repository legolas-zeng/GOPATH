//client.go
package main

import (
    "log"
    "net"
    "time"
    "strconv"
    "fmt"
    "encoding/json"
)

func send(conn net.Conn) {
    for i := 0; i < 10; i++ {
        time.Sleep(2 * time.Second)
        dic := make(map[string]interface{})
        dic["index"] = i
        dic["timestamp"] = time.Now().Format(time.RFC3339)
        jsonString, err := json.Marshal(dic)
        if err != nil {
            log.Println(err)
        }
        length := len(jsonString)
        if length > 99999 {
            //Header中标识字符串长度的最大为99999
            panic("data is too long to send")
        }
        lengthText := strconv.Itoa(length)
        textLength := fmt.Sprintf("%05s", lengthText)[:5]
        headerText := append([]byte("Header"), textLength...)
        jsonString = append(headerText, jsonString...)
        _, err = conn.Write([]byte(jsonString))
        if err != nil {
            log.Println(err)
        }
        log.Println("send : ", jsonString)
    }
    log.Println("send finished")
}

func main()  {
    server := "127.0.0.1:2048"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
    if err != nil {
        log.Println(err.Error())
        return
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        log.Println(err.Error())
        return
    }

    log.Println("connect success")
    send(conn)
    conn.Close()
}