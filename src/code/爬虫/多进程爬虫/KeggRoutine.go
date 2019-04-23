package main

import (
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Kegg struct {
	Genes 		string
	Organism 	string
	AAseq		string
	NTseq		string
}


func main(){
	t1 := time.Now()
	url := fmt.Sprintf("https://www.kegg.jp/dbget-bin/www_bget?ec:4.4.1.5")
	KegggetResponse(url)
	//fmt.Println(res)
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
	doc.Find("td.fr2 tbody").Each(func(i int, s *goquery.Selection) {
		txt:= s.Text()
		fmt.Println("==========",txt,"==========")
		GenesTr := s.Find("tr").Eq(1)
		fmt.Println(GenesTr)
		//img,_ :=s.Find("img").Attr("src")
		//num:=s.Find("em").Text()
		//star:=s.Find("span.rating_num").Text()
		//name,_:=s.Find("img").Attr("alt")
		//fmt.Println(img)
		//fmt.Println(num)
		//fmt.Println(star)
		//pages = append(pages, Kegg{
		//	Num: num,
		//	Url:  img,
		//	Star: star,
		//	Name: name,
		//})
	})
	return pages
}

