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
    conn, _ := NewRedisClient()
    type Data struct {
        Name *string
        Age *int
    }
    value,_ := json.Marshal(cmdinfo)
    fmt.Println(value)
    fmt.Printf("%T",value)
    _, err := conn.Do("Publish", "order-create", value)
    return err
}