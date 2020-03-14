package main

import (
	_ "RemoteCon/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"RemoteCon/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "zwa:qq1005521@tcp(192.168.3.5:3306)/pcinfo?charset=utf8", 30)
	orm.RegisterModel(
		new(models.Pcinfo),
	)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库创建失败!!")
		fmt.Println(err)
	} else {
		fmt.Printf("数据库初始化已完成！！")
	}

}
func indexdiv(index int) (index1 int) {
	index1 = index % 5
	return
}

func main() {
	beego.AddFuncMap("indexdiv", indexdiv)
	beego.SetStaticPath("/images","images")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")
	beego.Run()
}

