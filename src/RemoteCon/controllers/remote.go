package controllers

import (
    "github.com/astaxie/beego"
    "RemoteCon/models"
    "k8s.io/apimachinery/pkg/util/json"
    "fmt"
    "github.com/go-redis/redis"
)

type RemoteController struct {
    beego.Controller
}

var (
    cmdinfo=make(map[string]string)
)

func (this *RemoteController) GetPcinfo(){
    beego.ReadFromRequest(&this.Controller)

    PcInfo := &models.Pcinfo{}
    info:= PcInfo.GetPcInfo()
    this.Data["info"] = info
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "pcinfo/pcinfo.html"
    this.LayoutSections["js"] = "pcinfo/pcinfo_js.html"
    this.Render()
}


func (this *RemoteController) Remote(){
    beego.ReadFromRequest(&this.Controller)
    PcInfo := &models.Pcinfo{}
    info:= PcInfo.GetPcInfo()
    this.Data["info"] = info
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "remote/remote.html"
    this.LayoutSections["js"] = "remote/remote_js.html"
    this.Render()
}

func (this *RemoteController) ApiRemote(){
    ip := this.GetString("ip_list")
    cmds := this.GetString("cmd")
    argv1 := this.GetString("argv1")
    cmdinfo[ip]=argv1
    fmt.Println("+++++++",cmdinfo)
    resp := make(map[string]interface{})
    if cmds=="shell"{
        fmt.Println("往redis消息队列发送信息：",cmdinfo)
        conn, _ := NewRedisClient()
        value,_ := json.Marshal(cmdinfo)
        _, err := conn.Do("Publish", "order-create", value)
        if err != nil{
            resp["status"] = 0
            resp["msg"] = "刷新数据成功"
            this.Data["json"] = resp
        }else {
            resp["status"] = 1
            resp["msg"] = "刷新数据失败"
            this.Data["json"] = resp
        }
    }
    go subscribeReq()
    this.ServeJSON()
}


// 订阅返回的信息
func subscribeReq(){
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379", // redis地址
        Password: "", // redis密码，没有则留空
        DB:       0,  // 默认数据库，默认是0
    })
    sub := client.Subscribe("result")
    fmt.Println(sub.Receive())
    for msg := range sub.Channel(){
        fmt.Println(msg.Payload)

    }
    return
}



//func sendCmdToRedis(iplist map[int]string,shell string) error {
//    conn, _ := NewRedisClient()
//
//    value,_ := json.Marshal(cmdinfo)
//    fmt.Println(value)
//    fmt.Printf("%T",value)
//    _, err := conn.Do("Publish", "order-create", value)
//    return err
//}


