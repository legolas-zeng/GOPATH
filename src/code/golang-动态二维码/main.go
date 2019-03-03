package main

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	"image/gif"
	"github.com/docker/docker/image"
	"os"
)

type Water struct {
	image  image.Image
	gifImg *gif.GIF
}


func main() {
	err := qrcode.WriteFile("https://u.wechat.com/ELD9GwVqKqlEyGHiH4u0d3k", qrcode.Medium, 441, "E:\\GOPATH\\src\\code\\img\\qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}

func water(path string)(*Water,error){
	// E:\\GOPATH\\src\\code\\img\\candy.gif
	// E:\\GOPATH\\src\\code\\img\\qr.png
	// E:\\GOPATH\\src\\code\\img\\qr.gif
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var img image.Image
	var gifImg *gif.GIF
	gifImg, err = gif.DecodeAll(f)
	img = gifImg.Image[0]
	return &Water{
		image:   img,
		gifImg:  gifImg,

	}, nil
}

