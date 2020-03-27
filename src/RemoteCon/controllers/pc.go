package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "k8s.io/apimachinery/pkg/util/json"
)

type PcController struct {
    beego.Controller
}

func (this *PcController)ApiFlushPcInfo()  {
    ip := this.GetString("ip")
    cmd := this.GetString("cmd")
    fmt.Println(ip,cmd)
    cmdinfo := make(map[string]string)
    cmdinfo[ip] = cmd
    fmt.Println("+++++++",cmdinfo)
    err := publish(cmdinfo)
    resp := make(map[string]interface{})
    if err != nil {
        resp["status"] = 1
        resp["msg"] = "刷新数据错误"
    }else {
        resp["status"] = 0
        resp["msg"] = "刷新数据成功"
    }
    this.Data["json"] = resp
    this.ServeJSON()

}

func NewRedisClient() (conn redis.Conn, err error) {
    host := "127.0.0.1"
    port := "6379"
    adderss := host + ":" + port
    c, err := redis.Dial("tcp", adderss)
    return c, err
}

func publish(cmdinfo map[string]string)  error {
    fmt.Println("往redis消息队列发送信息：",cmdinfo)
    conn, _ := NewRedisClient()
    value,_ := json.Marshal(cmdinfo)
    _, err := conn.Do("Publish", "result", value)
    return err
}
