package routers

import (
	"RemoteCon/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.IndexController{},"get:Index")
	beego.Router("/remote",&controllers.RemoteController{},"*:Remote")

}
