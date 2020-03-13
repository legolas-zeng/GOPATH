package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
)

type PcController struct {
    beego.Controller
}

func (this *PcController)ApiFlushPcInfo()  {
    ip := this.GetString("ip")
    fmt.Println(ip)

}