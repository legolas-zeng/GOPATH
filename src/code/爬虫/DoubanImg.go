package main

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"net/http"
	"io"
	"fmt"
	"math/rand"
)

type Movie struct {
	Num 	string
	Url  	string
	Star 	string
	Name 	string

}

const (
	baseurl string = "https://movie.douban.com/top250?start=25&filter="
	imgpath string = "C:\\Users\\Administrator.000\\Desktop\\images"
)

func main(){
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100)) //返回[0,100)的随机整数
	}
	//res := getResponse(baseurl)
	//DownloadImg(res)
}

// 获取分页
func getResponse(url string)  []Movie{
	//func getResponse(url string)  *goquery.Document{
	content,err:= goquery.NewDocument(url)
	if err != nil{
		panic(err)
	}
	//fmt.Println(content.Html())
	return ParseResponse(content)
	//return content
}

func ParseResponse(doc *goquery.Document) (pages []Movie) {
	doc.Find("div.item").Each(func(i int, s *goquery.Selection) {
		img,_ :=s.Find("img").Attr("src")
		num:=s.Find("em").Text()
		star:=s.Find("span.rating_num").Text()
		name,_:=s.Find("img").Attr("alt")
		//fmt.Println(img)
		//fmt.Println(num)
		//fmt.Println(star)
		pages = append(pages, Movie{
			Num: num,
			Url:  img,
			Star: star,
			Name: name,
		})
	})
	return pages
}

func DownloadImg(pages []Movie){
	for _,v:= range pages{
		GetImg(v.Url,v.Name)
	}

}

func GetImg(url string , name string) {
	res, _ := http.Get(url)
	fmt.Println("save dir:", imgpath)
	file_name := imgpath + "\\" + name + ".jpg"
	fmt.Println("file:", file_name)
	file, _ := os.Create(file_name)
	io.Copy(file, res.Body)
}


