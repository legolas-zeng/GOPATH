package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BeginController struct {
	beego.Controller
}

func (this *BeginController) Function() {
	//this.TplName = "index.html"
	filename := this.GetString(":excelname")
	fmt.Println("获取到文件名：",filename)
	this.TplName = "function.html"
	this.Render()
}