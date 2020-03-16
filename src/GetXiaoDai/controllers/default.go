package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


type TestController struct{
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


func (this *TestController) Index(){
	paramMap := this.Ctx.Input.Params()
	//获取RESTFUL风格的参数
	//此时的URL为 localhost:8080/test/index/aaa/bbb
	this.Data["Website"] = paramMap["0"]        //aaa
	this.Data["Email"] = paramMap["1"]        //bbb

	v := this.GetSession("uid")
	this.Data["uid"] = v.(string)

	this.TplName = "test.html"
}
