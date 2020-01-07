package controllers

import "github.com/astaxie/beego"

type OrderController struct {
	beego.Controller
}


func (c *OrderController) Get() {
	c.Layout = "index.html"
	c.TplName = "order/order.html"
	c.Render()
}