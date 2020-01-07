package routers

import (
	"Garage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/order", &controllers.OrderController{})
}
