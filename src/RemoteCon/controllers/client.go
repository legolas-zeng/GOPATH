package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
)

type ClientController struct {
    beego.Controller
}

func (c *ClientController) Conn() {
    fmt.Println("111111")
}




