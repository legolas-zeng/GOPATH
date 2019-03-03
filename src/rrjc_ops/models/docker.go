package models

import "github.com/astaxie/beego/orm"

type Dockers struct {
	Id			int			`orm:"column(uid);pk;auto"` // 设置主键
	Group 		string 		`valid:"Required";orm:"size(10)"`
	Container 	string 		`valid:"Required";orm:"size(30)"`
	IpDocker	string		`valid:"Required";orm:"size(16)"`
	Node 		*Nodes		`orm:"rel(fk)"`  			//设置外键
	Status 		int64 		`orm:"default(0)"`
}
type Nodes struct {
	Id			int			`orm:"column(uid);auto"`
	Ip 			string 		`valid:"Required";orm:"size(16);unique;pk"`
	Status 		int64 		`orm:"default(0)";orm:"size(1)"`
	HostName 	string 		`orm:"default(server-00)"`
}

type Groups struct {
	Id          int64    		`pk:"auto" form:"id"`
	ProName 	string
	Ip 			string
	GroupName	string		`orm:"default(group-0)"`
	Status		int64		`orm:"default(0)"`

}

//TODO dokcer信息查询
//func (c *Dockers) FindAllDockInfo() ([]*Dockers, error) {
//
//	var dock []*Dockers
//	o := orm.NewOrm()
//	qs := o.QueryTable("dockers")
//	qs = qs.Limit(100)
//	_, err := qs.OrderBy("group").All(&dock)
//
//	//common.PanicIf(err)
//	return dock, err
//
//}
//TODO 联表查询
func (c *Dockers) FindAllDockInfo() ([]*Dockers) {

	var dock []*Dockers
	o := orm.NewOrm()
	o.QueryTable("dockers").RelatedSel().All(&dock)

	//common.PanicIf(err)
	return dock

}

//TODO 查询信息
func (c *Dockers) FindDockerInfo(table string,filer string,) ([]*Dockers) {
	var dock []*Dockers
	o := orm.NewOrm()
	o.QueryTable(table).Filter("IpDocker",filer).All(&dock,"Node")
	return dock
}

//TODO 更新数据表




//TODO 保存数据表
func (this *Groups) SaveGroupInfo() error {

	_, err := orm.NewOrm().Insert(this)

	return err
}


