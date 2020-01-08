package main

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"net/http"
	"io"
	"fmt"
	"time"
	"sync"
)

type Movie struct {
	Url  	string
	Name 	string

}

const (
	//baseurl string = "https://movie.douban.com/top250?start=25&filter="
	imgpath string = "C:\\Users\\Administrator\\Desktop\\images1\\Uncensored"
)

var waitgroup sync.WaitGroup

func main(){
	t1 := time.Now()
	for i := 722; i < 1722; i++ {
		//url := fmt.Sprintf("http://javgod.net/?paged=%v&cat=23", i)
		url := fmt.Sprintf("http://filmhav.com/category/uncensored/page/%v",i)
		fmt.Printf("整在爬取第%v页",i)
		res := getResponse(url)
		fmt.Println(res)
		waitgroup.Add(1) //计数器+1 可以认为是队列+1
		go DownloadImg(res)
	}
	waitgroup.Wait() //进行阻塞等待 如果 队列不跑完 一直不终止
	elapsed := time.Since(t1)
	fmt.Println("总共用时: ", elapsed)
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
	doc.Find("div.type-post").Each(func(i int, s *goquery.Selection) {
		img,_ :=s.Find("img").Attr("src")
		name:=s.Find("h2").Text()
		pages = append(pages, Movie{
			Url:  img,
			Name: name,
		})
	})
	return pages
}

//func ParseResponse(doc *goquery.Document) (pages []Movie) {
//	doc.Find("article").Each(func(i int, s *goquery.Selection) {
//		img,_ :=s.Find("img").Attr("src")
//		name:=s.Find("h2").Text()
//		pages = append(pages, Movie{
//			Url:  img,
//			Name: name,
//		})
//	})
//	return pages
//}

func DownloadImg(pages []Movie){
	for _,v:= range pages{
		GetImg(v.Url,v.Name)
	}
	defer waitgroup.Done() //如果跑完就进行 队列-1

}

func GetImg(url string , name string) {
	res, err := http.Get(url)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("A error occurred!")
	}else {
		file_name := imgpath + "\\" + name + ".jpeg"
		file, _ := os.Create(file_name)
		io.Copy(file, res.Body)
	}
}


