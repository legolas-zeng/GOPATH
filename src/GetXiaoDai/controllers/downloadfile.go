package controllers

import "github.com/astaxie/beego"

type FileOptDownloadController struct {
	beego.Controller
}
func (this *FileOptDownloadController) DownloadFile() {
	//图片,text,pdf文件全部在浏览器中显示了，并没有完全的实现下载的功能
	//this.Redirect("/static/img/1.jpg", 302)

	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	this.Ctx.Output.Download("static/img/a1.jpg","a1.jpg")
}