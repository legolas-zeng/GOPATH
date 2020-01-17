package main

import (
    _ "GetXiaoDai/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "fmt"
    "GetXiaoDai/models"
)

func init() {
    orm.Debug = true
    orm.RegisterModel(new(models.XiaoDai))
    orm.RegisterDataBase("default", "mysql", "root:qq1005521@tcp(127.0.0.1:3306)/xiaodai?charset=utf8", 30)
    err := orm.RunSyncdb("default", false, true)
    if err != nil {
        fmt.Println("数据库创建失败!!")
        fmt.Println(err)
    } else {
        fmt.Printf("数据库初始化已完成！！")
    }
}

func main() {
    beego.SetStaticPath("/static/images", "images")
    beego.SetStaticPath("/static/css", "css")
    beego.SetStaticPath("/static/js", "js")
    orm.Debug = true

    beego.Run()
}
