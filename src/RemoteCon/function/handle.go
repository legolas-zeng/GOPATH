package function

import (
    "github.com/astaxie/beego"
    "time"
    "net"
    "sync"
    "github.com/mxi4oyu/MoonSocket/protocol"
)


func ReadFromSocketConn() {
    maps := GetSocketMaps()
    for key, value := range maps {
        tmpbuf:=make([] byte,0)
        buf:=make([] byte,1024)
        conn := *value
        length, err := conn.Read(buf)
        if err != nil {
            continue
        }
        tmpbuf = protocol.Depack(append(tmpbuf,buf[:length]...))
        beego.Info("接收到来自 : ", key, " 的信息: ", string(tmpbuf))
    }
}

func WriteToSocketConn(b []byte) {
    maps := GetSocketMaps()
    for key, value := range maps {
        conn := *value
        smsg:=protocol.Enpack(b)
        conn.Write(smsg)
        beego.Info("发送信息给: ", key, " content: ", smsg)
    }
}


func LoopReadFromSocketConn() {
    for {
        ReadFromSocketConn()
    }
}

func LoopWriteToSocketConn() {
    var sendStr = "服务端信息"
    for {
        WriteToSocketConn([]byte(sendStr))
        Sleep5S()
    }
}

func Sleep5S() {
    time.Sleep(time.Duration(5) * time.Second)
}

var socketMaps map[string]*net.Conn

var once sync.Once

func GetSocketMaps() map[string]*net.Conn {
    once.Do(func() {
        socketMaps = make(map[string]*net.Conn)
    })
    return socketMaps
}