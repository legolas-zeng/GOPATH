package main

import (
	_ "Garage/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Garage/models"
)

func init(){
	models.RegisterDB()
}
func main() {
	beego.SetStaticPath("/static/images", "images")
	beego.SetStaticPath("/static/css", "css")
	beego.SetStaticPath("/static/js", "js")
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

