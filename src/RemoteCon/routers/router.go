package routers

import (
	"RemoteCon/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	cli := beego.NewNamespace("/client",
		beego.NSRouter("/",&controllers.ClientController{},"*:Conn"),
		beego.NSInclude(
			&controllers.ClientController{},

		),
	)


	beego.AddNamespace(cli)
}
