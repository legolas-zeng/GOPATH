package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type XiaoDai struct {
	Id			int
	FileName 	string		//文件名
	HashName	string		//哈希名
	FilePath	string		//文件路径
	Path		string		//提取的数据路径

}


//TODO 查询信息
func (this *XiaoDai) FindXiaoDaiInfo(table string,filer string,) ([]*XiaoDai) {
	var xiaodai []*XiaoDai
	o := orm.NewOrm()
	o.QueryTable(table).Filter("IpDocker",filer).All(&xiaodai,"Node")
	return xiaodai
}

//TODO 保存数据表
func (this *XiaoDai) InsertXiaoDaiInfo(FileName string,HashName string,FilePath string,Path string) {
	var xiaodai XiaoDai
	o := orm.NewOrm()
	xiaodai.FileName = FileName
	xiaodai.HashName = HashName
	xiaodai.FilePath = FilePath
	xiaodai.Path = Path
	id ,err := o.Insert(&xiaodai)
	if err == nil {
		fmt.Println(id)
	}
}

//TODO 更新数据表




