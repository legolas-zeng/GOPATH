package routers

import (
	"rrjc_ops/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test",&controllers.TestController{})

	dok := beego.NewNamespace("/docker",
		beego.NSRouter("/",&controllers.DockerListController{},"*:DockerList"),
		beego.NSRouter("/add",&controllers.DockerListController{},"*:DockerAdd"),
		beego.NSRouter("/function",&controllers.DockerListController{},"*:DockerFunction"),

		beego.NSInclude(
			&controllers.DockerListController{},

		),
	)


	beego.AddNamespace(dok)
}
