package main

import (
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type info struct {
	url 		string
	Organism 	string
	AAseq		string
	NTseq		string
} 

func main(){
	t1 := time.Now()
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", "root", "qq1005521", "tcp", "192.168.3.5", 3306, "gosql")
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("打开数据库失败,err:%v\n", err)
		return
	}
	var urllist []string
	list := queryMulti(DB,urllist)
	hendleUrl(list)

	elapsed := time.Since(t1)
	fmt.Println("总共用时: ", elapsed)
}


func queryMulti(DB *sql.DB,urllist []string) (url []string){
	rows, err := DB.Query("select url from kegg limit ?", 10)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			fmt.Printf("Scan failed,err:%v", err)
		}
		urllist = append(urllist,url)
	}
	return urllist
}
//www.kegg.jp/dbget-bin/www_bget?loki:Lokiarch_02820
func hendleUrl(urllist []string)  {
	for _,v:= range urllist{
		fullurl := strings.Join([]string{"https://www.kegg.jp", v}, "")
		kegggetResponse(fullurl,v)
	}
}

func kegggetResponse(url string,v string) []info{
	content,err:= goquery.NewDocument(url)
	if err != nil{
		panic(err)
	}
	//fmt.Println(content.Html())
	return keggParseResponse(content,v)

}
func keggParseResponse(doc *goquery.Document,v string) (data []info)  {
	doc.Find("body > div > table > tbody > tr > td > table:nth-child(2) > tbody > tr > td > form > table > tbody").Each(func(i int, s *goquery.Selection) {
		str := s.Find("tr:nth-child(5) > td > div").Text()
		AAseq := s.Find("tr:nth-child(13) > td").Text()
		NTseq := s.Find("tr:nth-child(14) > td").Text()
		fmt.Println("==========",NTseq,"==========")
		data = append(data, info{
			url		 :  v,
			Organism :	str,
			AAseq	 :	AAseq,
			NTseq	 :	NTseq,
		})

	})
	return data
}