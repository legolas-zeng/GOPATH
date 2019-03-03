package main

import (
	"fmt"
	"github.com/opesun/goquery"
)

const (
	baseurl string = "https://movie.douban.com/top250"
)

func main(){
	response := getResponse(baseurl)
	fmt.Println(response)
}

func getResponse(url string)  interface{}{
	content,err:=goquery.ParseUrl(url)
	if err != nil{
		panic(err)
	}
	fmt.Println(content.Html())
	return content
}
