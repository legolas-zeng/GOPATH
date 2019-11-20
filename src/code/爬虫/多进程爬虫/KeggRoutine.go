package main

import (
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type Kegg struct {
	Genes 		string
}

const (
	USERNAME = "root"
	PASSWORD = "qq1005521"
	NETWORK  = "tcp"
	SERVER   = "192.168.3.5"
	PORT     = 3306
	DATABASE = "gosql"
)


func main(){
	t1 := time.Now()
	url := fmt.Sprintf("https://www.kegg.jp/dbget-bin/www_bget?ec:4.4.1.5")
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	res := KegggetResponse(url)
	insertData(DB,res)
	elapsed := time.Since(t1)
	fmt.Println("总共用时: ", elapsed)
}


// 获取分页
func KegggetResponse(url string) []Kegg{
	content,err:= goquery.NewDocument(url)
	if err != nil{
		panic(err)
	}
	//fmt.Println(content.Html())
	return KeggParseResponse(content)

}
func KeggParseResponse(doc *goquery.Document) (pages []Kegg) {
	doc.Find("td.fr2 > table > tbody > tr:nth-child(13) > td > div > table > tbody > tr > td:nth-child(2) > a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		//fmt.Println("==========",url,"==========")
		pages = append(pages, Kegg{
			Genes:url,
		})
	})
	return pages
}
// url格式：https://www.kegg.jp/dbget-bin/www_bget?loki:Lokiarch_02820
//func HandleUrl(pages []Kegg){
//	for _,v:= range pages{
//		fmt.Println(v.Genes)
//	}
//}

//插入数据
func insertData(DB *sql.DB,pages []Kegg){
	for _,v:= range pages {
		result, err := DB.Exec("insert INTO kegg(url) values(?)", v.Genes)
		if err != nil {
			fmt.Printf("插入数据失败,错误详情:%v", err)
			return
		}
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			fmt.Printf("最后插入ID,err:%v", err)
			return
		}
		fmt.Println("LastInsertID:", lastInsertID)
		rowsaffected, err := result.RowsAffected()
		if err != nil {
			fmt.Printf("受影响的行数,err:%v", err)
			return
		}
		fmt.Println("RowsAffected:", rowsaffected)
	}
}

