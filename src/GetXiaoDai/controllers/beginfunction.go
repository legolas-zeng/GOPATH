package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"path/filepath"
	"os"
)

type BeginController struct {
	beego.Controller
}

var (
	DesPath = "C:\\Users\\Administrator\\Desktop"
	sema = make(chan struct{}, 20)
	)

func (this *BeginController) Function() {
	//this.TplName = "index.html"
	filename := this.GetString(":excelname")
	fmt.Println("获取到文件名：",filename)
	this.TplName = "function.html"
	this.Render()
}

func (this *BeginController) GetFileSize (){
	var size int64
	err := filepath.Walk(DesPath, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		fmt.Println("文件夹大小读取失败！")
	}else {
		size := fmt.Sprintf("%.1f MB", float64(size)/1e6)
		result := struct {
			Val string
		}{size}
		this.Data["json"] = &result
		this.ServeJSON()
	}
}




