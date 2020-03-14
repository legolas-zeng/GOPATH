package models

import "github.com/astaxie/beego/orm"

type Pcinfo struct {
    Id			int			`orm:"size(11);auto"`
    Ip 		    string 		`orm:"size(255):null"`
    Cpu 	    string 		`orm:"size(255):null"`
    Osinfo	    string		`orm:"size(255):null"`
    Men 		string		`orm:"size(255):null"`
    Online 		string 		`orm:"size(255):null"`
    Port 		string 		`orm:"size(255):null"`
    Pcname 		string 		`orm:"size(255):null"`
}

//TODO 查询信息
func (c *Pcinfo) GetPcInfo() ([]*Pcinfo) {
   var pc []*Pcinfo
   o := orm.NewOrm()
   o.QueryTable("pcinfo").RelatedSel().All(&pc)
   return pc
}
