package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type xiaodai struct {
	Id			int
	Name 		string		//客户名称
	Money		int			//消费金额
	Close		bool		//是否关闭

}

func RegisterDB(){
	orm.Debug = true
	orm.RegisterModel(new(xiaodai))
	orm.RegisterDataBase("default", "mysql", "root:qq1005521@tcp(127.0.0.1:3306)/xiaodai?charset=utf8", 30)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库创建失败!!")
		fmt.Println(err)
	} else {
		fmt.Printf("数据库初始化已完成！！")
	}

}
