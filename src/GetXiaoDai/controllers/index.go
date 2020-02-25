package controllers

import (
    "github.com/astaxie/beego"
    "os"
    "fmt"
)

type IndexController struct {
    beego.Controller
}


func (this *IndexController) Index() {
    //this.TplName = "index.html"
    DeleteFile(ExeclFile)
    DeleteFile(LogFile)
    this.TplName = "form_file_upload.html"
    this.Render()
}

func DeleteFile(filename string){
    err := os.Remove(filename)
    if err != nil {
        fmt.Printf("删除文件失败：%s \n",err)
    } else {
        fmt.Printf("删除文件%s成功 \n",filename)
    }
}

