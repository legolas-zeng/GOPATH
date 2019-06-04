package main

import "fmt"

func double(x *int) {
	*x += *x //*x = *x + *x 翻倍
	x = nil //拷贝传递x的值，对x不影响
	fmt.Println(x)
}

func main() {
	var a = 3
	double(&a) //&取地址
	fmt.Println(a) // 6

	p := &a
	double(p)
	fmt.Println(a, p == nil)
}

//限制一：Go 的指针不能进行数学运算。
//限制二：不同类型的指针不能相互转换。
//限制三：不同类型的指针不能使用 == 或 != 比较。
//限制四：不同类型的指针变量不能相互赋值。