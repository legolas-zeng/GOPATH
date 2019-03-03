package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *TestController) Get() {
	c.TplName = "test/test.html"
	c.Layout = "test/test.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["content"] = "test/content.html"
}
