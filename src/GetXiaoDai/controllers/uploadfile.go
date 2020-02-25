package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
    "path"
    "os"
)

type UploadFileController struct {
    //BaseController    //这个是自己封装的controller
    beego.Controller
}


func (this *UploadFileController) UpFile(){

    f, h, _ := this.GetFile("myfile")//获取上传的文件
    fmt.Println(h.Filename)
    //获取后缀名
    ext := path.Ext(h.Filename)
    //验证后缀名是否符合要求
    var AllowExtMap map[string]bool = map[string]bool{
        ".jpg":true,
        ".jpeg":true,
        ".png":true,
        ".zip":true,
        ".xlsx":true,
    }
    if _,ok:=AllowExtMap[ext];!ok{
        this.Ctx.WriteString( "后缀名不符合上传要求" )
        return
    }
    //创建目录
    //uploadDir := DesPath + time.Now().Format("2006/01/")
    err := os.MkdirAll( DesPath , 777)
    if err != nil {
        this.Ctx.WriteString( fmt.Sprintf("%v",err) )
        return
    }
    //构造文件名称
    //rand.Seed(time.Now().UnixNano())
    //hashName := time.Now().Format("2006-01-02-15-04-05")
    //fmt.Println(hashName)
    //fmt.Printf("%T",*str)
    //fmt.Println("1111111",reflect.TypeOf(hashName))
    fpath := DesPath + "\\" + h.Filename
    fmt.Println(fpath)
    defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
    err = this.SaveToFile("myfile", fpath)
    if err != nil {
        this.Ctx.WriteString( fmt.Sprintf("%v",err) )
    }
    //把文件名返回给前端
    this.Ctx.WriteString( h.Filename )

    //var XiaoDai models.XiaoDai
    //XiaoDai.InsertXiaoDaiInfo(fileName,hashName,fpath,fpath)
}

