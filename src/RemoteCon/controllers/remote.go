package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
)

type RemoteController struct {
    beego.Controller
}

func (this *RemoteController) Remote(){
    fmt.Println("-----")
    beego.ReadFromRequest(&this.Controller)
    this.TplName = "base/base.html"
    this.Layout = "base/base.html"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["re_content"] = "remote/pcinfo.html"
    this.Render()
}
