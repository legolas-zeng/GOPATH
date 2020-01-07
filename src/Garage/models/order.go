package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	Id			int
	Name 		string		//客户名称
	CarID		string		//车牌号
	CardType	string		//车型
	Phone		string		//客服手机号
	InDate		string		//入店时间
	OutDate		string		//离店时间
	Goods		string		//消费商品
	Money		int			//消费金额
	Close		bool		//是否关闭
	Ticket  	bool		//是否开票
}


func RegisterDB(){
	orm.Debug = true
	orm.RegisterModel(new(Order))
	orm.RegisterDataBase("default", "mysql", "zwa:qq1005521@tcp(192.168.3.5:3306)/garage?charset=utf8", 30)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库创建失败!!")
		fmt.Println(err)
	} else {
		fmt.Printf("数据库初始化已完成！！")
	}

}

