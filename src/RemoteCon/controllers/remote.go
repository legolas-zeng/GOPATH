package controllers

import (
    "github.com/astaxie/beego"
    "RemoteCon/models"
)

type RemoteController struct {
    beego.Controller
}

func (this *RemoteController) Remote(){
    beego.ReadFromRequest(&this.Controller)

    PcInfo := &models.Pcinfo{}
    info:= PcInfo.GetPcInfo()
    this.Data["info"] = info
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "remote/pcinfo.html"
    this.LayoutSections["js"] = "remote/pcinfo_js.html"
    this.Render()
}


