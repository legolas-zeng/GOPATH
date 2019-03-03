package main

import (
	_ "rrjc_ops/routers"
	"github.com/astaxie/beego"

	"fmt"
	"github.com/astaxie/beego/orm"
	"rrjc_ops/models"
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:qq1005521@tcp(127.0.0.1:3306)/ops?charset=utf8", 30)
	orm.RegisterModel(
		new(models.Dockers),
		new(models.Nodes),
		new(models.Groups),
		)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库创建失败!!")
		fmt.Println(err)
	} else {
		fmt.Printf("数据库初始化已完成！！")
	}

}

func main() {
	beego.SetStaticPath("/images","images")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")
	beego.Run()
}

