package main

import (
    _ "RemoteCon/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "RemoteCon/models"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
    "net"
    "flag"
    "RemoteCon/function"
    "strings"
)

var (
    inputIP = flag.String("IP", "0.0.0.0", "Listen IP")
)

func init() {
    orm.Debug = true
    orm.RegisterDataBase("default", "mysql", "zwa:qq1005521@tcp(192.168.3.5:3306)/pcinfo?charset=utf8", 30)
    orm.RegisterModel(
        new(models.Pcinfo),
    )
    err := orm.RunSyncdb("default", false, true)
    if err != nil {
        beego.Error("数据库创建失败!!")
        beego.Error(err)
    } else {
        beego.Info("数据库初始化已完成！！")
    }

}
func indexdiv(index int) (index1 int) {
    index1 = index % 5
    return
}

func main() {
    //添加模板方法
    beego.AddFuncMap("indexdiv", indexdiv)
    beego.SetStaticPath("/images", "images")
    beego.SetStaticPath("/css", "css")
    beego.SetStaticPath("/js", "js")
    //go tcpSocketListen()
    //go function.LoopReadFromSocketConn()
    //go function.LoopWriteToSocketConn()
    beego.Run()
    //go controllers.SubscribeReq()

}

func tcpSocketListen() {
    listenPort, parseErr := beego.AppConfig.Int("socketlistenport")
    if parseErr != nil {
        beego.Error("socket端口错误 :", beego.AppConfig.String("socketlistenport"))
        return
    }

    listener, listenErr := net.Listen("tcp", *inputIP+":"+strconv.Itoa(listenPort))
    if listenErr != nil {
        beego.Error("socket监听错误 !")
        return
    } else {
        beego.Info("socket监听成功：")
    }

    for {
        conn, acceptErr := listener.Accept()
        if acceptErr != nil {
            continue
        }

        maps := function.GetSocketMaps()
        add := strings.Split(conn.RemoteAddr().String(),":")
        maps[add[0]] = &conn
        beego.Debug(conn)
    }
}
