package routers

import (
	"RemoteCon/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.RemoteController{},"*:GetPcinfo")
	remote := beego.NewNamespace("/remote",
		beego.NSRouter("/",&controllers.RemoteController{},"*:Remote"),
		beego.NSRouter("/api",&controllers.RemoteController{},"*:ApiRemote"),
		beego.NSInclude(
			&controllers.RemoteController{},
		),
	)
	pc := beego.NewNamespace("/pc",
		beego.NSRouter("/reflush",&controllers.PcController{},"*:ApiFlushPcInfo"),

		beego.NSInclude(
			&controllers.PcController{},
		),
	)

	beego.AddNamespace(pc)
	beego.AddNamespace(remote)

}
