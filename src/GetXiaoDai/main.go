package main

import (
	_ "GetXiaoDai/routers"
	"github.com/astaxie/beego"
	"GetXiaoDai/models"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}

