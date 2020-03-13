package routers

import (
	"RemoteCon/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.IndexController{},"get:Index")
	beego.Router("/remote",&controllers.RemoteController{},"*:Remote")

	pc := beego.NewNamespace("/pc",
		beego.NSRouter("/reflush",&controllers.PcController{},"*:ApiFlushPcInfo"),

		beego.NSInclude(
			&controllers.PcController{},
		),
	)

	beego.AddNamespace(pc)

}
