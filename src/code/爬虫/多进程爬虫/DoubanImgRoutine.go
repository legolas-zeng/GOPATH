package 多进程爬虫

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
	Num 	string
	Url  	string
	Star 	string
	Name 	string

}

const (
	//baseurl string = "https://movie.douban.com/top250?start=25&filter="
	imgpath string = "C:\\Users\\Administrator\\Desktop\\images"
)

var waitgroup sync.WaitGroup

func main(){
	t1 := time.Now()
	for i := 0; i < 11; i++ {
		url := fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", i*25)
		fmt.Printf("整在爬取第%v页",i+1)
		res := getResponse(url)
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
	defer waitgroup.Done() //如果跑完就进行 队列-1

}

func GetImg(url string , name string) {
	res, _ := http.Get(url)
	file_name := imgpath + "\\" + name + ".jpg"
	file, _ := os.Create(file_name)
	io.Copy(file, res.Body)
}


