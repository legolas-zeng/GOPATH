package main

import (
    "fmt"
    "net"
    "time"
    "bufio"
    "io"
    "strings"
    "strconv"
    "os"
)

const (
    IP = "*.*.*.*:1530"
)

func main() {
    var tcpAddr *net.TCPAddr
    tcpAddr,_ = net.ResolveTCPAddr("tcp","127.0.0.1:8082")

    conn,err := net.DialTCP("tcp",nil,tcpAddr)

    if err!=nil {
        fmt.Println("服务器连接失败 ! " + err.Error())
        return
    }

    defer conn.Close()

    fmt.Println(conn.LocalAddr().String() + " : 服务端已连接!")
    //onMessageReceived(conn)


}

func onMessageReceived(conn *net.TCPConn) {

    reader := bufio.NewReader(conn)
    b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
    conn.Write(b)
    for {
        msg, err := reader.ReadString('\n')
        fmt.Println("ReadString")
        fmt.Println(msg)

        if err != nil || err == io.EOF {
            fmt.Println(err)
            break
        }
        time.Sleep(time.Second * 2)

        fmt.Println("writing...")

        b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
        _, err = conn.Write(b)

        if err != nil {
            fmt.Println(err)
            break
        }
    }
}

func handleClient(conn net.Conn) {
    conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
    request := make([]byte, 128) // set maxium request length to 128B to prevent flood attack
    defer conn.Close()  // close connection before exit
    for {
        read_len, err := conn.Read(request)

        if err != nil {
            fmt.Println(err)
            break
        }

        if read_len == 0 {
            break // connection already closed by client
        } else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
            daytime := strconv.FormatInt(time.Now().Unix(), 10)
            conn.Write([]byte(daytime))
        } else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }

        request = make([]byte, 128) // clear last read content
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}