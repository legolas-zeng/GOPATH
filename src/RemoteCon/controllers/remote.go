package controllers

import (
    "github.com/astaxie/beego"
    "RemoteCon/models"
    "k8s.io/apimachinery/pkg/util/json"
    "github.com/go-redis/redis"
    "fmt"
    "regexp"
)

type RemoteController struct {
    beego.Controller
}

type Message struct {
    Channel string
    Pattern string
    Payload string
}

var (
    cmdinfo = make(map[string]string)
    reqinfo = make(map[string]string)
    datainfo = make(map[string]string)
)

func (this *RemoteController) GetPcinfo() {
    beego.ReadFromRequest(&this.Controller)

    PcInfo := &models.Pcinfo{}
    info := PcInfo.GetPcInfo()
    this.Data["info"] = info
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "pcinfo/pcinfo.html"
    this.LayoutSections["js"] = "pcinfo/pcinfo_js.html"
    this.Render()
}

func (this *RemoteController) Remote() {
    beego.ReadFromRequest(&this.Controller)
    PcInfo := &models.Pcinfo{}
    info := PcInfo.GetPcInfo()
    this.Data["info"] = info
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "remote/remote.html"
    this.LayoutSections["js"] = "remote/remote_js.html"
    this.Render()
}

func (this *RemoteController) ApiRemote() {
    ip := this.GetString("ip_list")
    cmds := this.GetString("cmd")
    argv1 := this.GetString("argv1")
    remoteip := handleData(ip)
    cmdinfo[ip] = argv1
    fmt.Println("+++++++", cmdinfo)
    //resp := make(map[string]interface{})
    if cmds == "shell" {
        fmt.Println("往redis消息队列发送信息：", cmdinfo)
        conn, _ := NewRedisClient()
        value, _ := json.Marshal(cmdinfo)
        _, err := conn.Do("Publish", "command", value)
        if err != nil {
            //resp["status"] = 0
            //resp["msg"] = "刷新数据成功"
            //this.Data["json"] = resp
            beego.Info("数据刷新成功")
        } else {
            //resp["status"] = 1
            //resp["msg"] = "刷新数据失败"
            //this.Data["json"] = resp
            beego.Error("数据刷新失败")
        }
    }
    //传递完消息后清空
    delete(cmdinfo, ip)
    //wg := &sync.WaitGroup{}
    //wg.Add(1)
    ////开始监听返回值
    //go subscribeReq(wg,remoteip)
    //wg.Wait()
    //fmt.Println("---------over----------")
    //this.ServeJSON()
    reqdata := subscribeReq(remoteip)
    this.Data["json"] = reqdata
    this.ServeJSON()
}

//用同步监听redis订阅信息
func subscribeReq(remoteip []string) map[string]string{
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379", // redis地址
        Password: "",               // redis密码，没有则留空
        DB:       0,                // 默认数据库，默认是0
    })
    fmt.Println("1", remoteip)
    for len(remoteip) !=0{
        redisSubscript := client.Subscribe("result")
        msg, err := redisSubscript.ReceiveMessage()
        if err != nil {
            client.Close()
        }
        json.Unmarshal([]byte(msg.Payload), &reqinfo)
        fmt.Println("========", reqinfo)
        for k, v := range reqinfo {
            removesli(&remoteip, k)
            datainfo[k] = v
        }
        fmt.Println("2", remoteip)
        //通过GC清空reqinfo
        reqinfo = make(map[string]string)
        continue
    }

    fmt.Println("+++++++++",datainfo)
    return datainfo
}

// 用异步订阅返回的信息
//func subscribeReq(wg *sync.WaitGroup,remoteip []string) {
//    client := redis.NewClient(&redis.Options{
//        Addr:     "127.0.0.1:6379", // redis地址
//        Password: "",               // redis密码，没有则留空
//        DB:       0,                // 默认数据库，默认是0
//    })
//    fmt.Println("1",remoteip)
//    for {
//        redisSubscript := client.Subscribe("result")
//        msg, err := redisSubscript.ReceiveMessage()
//        if err != nil {
//            client.Close()
//        }
//        json.Unmarshal([]byte(msg.Payload), &reqinfo)
//        fmt.Println("========",reqinfo)
//        for k, _ := range reqinfo {
//            removesli(&remoteip, k)
//            //wg.Done()
//        }
//        fmt.Println("2",remoteip)
//        if len(remoteip) == 0{
//            fmt.Println("退出")
//            wg.Done()
//            break
//
//        }
//        //wg.Done()
//        //for msg := range redisSubscript.Channel(){
//        //    fmt.Println("------",msg.Payload)
//        //}
//        }
//    return
//
//}

func removesli(remoteip *[]string, elem string) {
    for k, v := range *remoteip {
        if v == elem {
            kk := k + 1
            *remoteip = append((*remoteip)[:k], (*remoteip)[kk:]...)
        }
    }
}

func handleData(args string) []string {
    var ip_list []string
    reg1 := regexp.MustCompile(`(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d
)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`)
    ips := reg1.FindAll([]byte(args), -1)
    for _, value := range ips {
        ip_list = append(ip_list, string(value))
    }
    return ip_list
}
