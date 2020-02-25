package routers

import (
	"github.com/astaxie/beego"
	"GetXiaoDai/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/",&controllers.IndexController{},"get:Index")
	beego.Router("/uploadfile",&controllers.UploadFileController{},"*:UpFile")
	beego.Router("/download",&controllers.FileOptDownloadController{},"*:DownloadFile")
	beego.Router("/continue/:excelname",&controllers.BeginController{},"*:Function")
	beego.Router("/getfilesize",&controllers.BeginController{},"get:GetFileSize")
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.NSInclude(
		&controllers.UploadFileController{},
	)
}
