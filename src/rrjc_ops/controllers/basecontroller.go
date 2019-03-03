package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	isLogin bool
}
func (c *BaseController) Prepare() {
	userLogin := c.GetSession("userLogin")
	if userLogin == nil {
		c.isLogin = false
	} else {
		c.isLogin = true
	}
	c.Data["isLogin"] = c.isLogin
}
