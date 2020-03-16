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
    username:=this.GetSession("username")
    fmt.Println("+++++++",username)
    DeleteFile(ExeclFile)
    DeleteFile(LogFile)
    DeleteFolder(DesPaths)
    this.TplName = "form_file_upload.html"
    this.Data["username"] = username
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

func DeleteFolder(path string){
    err := os.RemoveAll(path)
    if err != nil {
        fmt.Printf("删除文件夹失败：%s \n",err)
    } else {
        fmt.Printf("删除文件夹%s成功 \n",path)
    }
}

