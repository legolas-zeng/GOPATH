package main

import (
	"github.com/Comdex/imgo"
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	img, err := imgo.DecodeImage("E:\\GOPATH\\src\\code\\img\\hsq.jpg") // 获取 图片 image.Image 对象
	if err != nil {
		fmt.Println(err)
	}
	height := imgo.GetImageHeight(img) // 获取 图片 高度[height]
	width := imgo.GetImageWidth(img)   // 获取 图片 宽度[width]
	imgMatrix := imgo.MustRead("E:\\GOPATH\\src\\code\\img\\hsq.jpg")

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetActiveSheet(index)

	for starty := 0; starty < height; starty++ {
		for startx := 0; startx < width; startx++ {
			C := imgMatrix[starty][startx][0:3]
			fmt.Println(C)
			var b string
			a := ""
			for i :=0; i<3; i++ {
				n := C[i]
				x := int64(n)
				b = dectohex(x)
				a += strings.Join([]string{b}, "") //拼接颜色代码
			}
			fmt.Println(a)
			style, err := xlsx.NewStyle(fmt.Sprintf(`{"fill":{"type":"pattern","color":["%s"],"pattern":1}}`,a))
			if err != nil {
				fmt.Println(err)
			}
			xx := startx+1
			yy := starty+1
			fmt.Printf("坐标点：%v,%v\n",xx,yy)
			coordinatex := GetLetterByNum(xx) //二十六进制转换
			cell := coordinatex + strconv.Itoa(yy) //cell需要x轴和y轴字符串相连，类似AA12，就代表第27列，12行
			fmt.Println(cell)
			xlsx.SetCellStyle("Sheet1", cell, cell, style)
			//R := imgMatrix[starty][startx][0] // 0,1,2,3,分别是R，G，B，A的值
		}
	}
	errs := xlsx.SaveAs("E:\\GOPATH\\src\\code\\img\\Workbook.xlsx")
	if err != nil {
		fmt.Println(errs)
		os.Exit(1)
	}

}

func dectohex(n int64) string {
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	if len(s) == 1{ //如果是一位十六进制的数，就在前面加个0
		s = "0" + s
	}
	return s
}

func GetLetterByNum(num int) (s string) {
	if num <= 0 {
		panic("num参数取值范围为大于零的整数")
	}
	var temp []rune
	yu := 0
	shang := num
	for {
		yu = shang % 26
		shang = shang / 26
		if yu == 0 {
			yu = 26
			shang--
		}
		temp = append(temp, rune(yu+'A'-1))
		if shang == 0 {
			break
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		s += string(temp[i])
	}
	return
}