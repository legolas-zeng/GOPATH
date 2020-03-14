package controllers

import (
    "github.com/astaxie/beego"
    "RemoteCon/models"
    "fmt"
)

type RemoteController struct {
    beego.Controller
}

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
    cmd := this.GetString("cmd")
    argv1 := this.GetString("argv1")
    fmt.Println(ip,cmd,argv1)
    resp := make(map[string]interface{})
    resp["status"] = 0
    resp["msg"] = "刷新数据成功"
    this.Data["json"] = resp
    this.ServeJSON()
}


