package main

import "fmt"

//148 131 115

//func main(){
//	var a []uint8
//	var b string
//	a = []uint8{148,131,15}
//	c := ""
//	for i :=0; i<3; i++ {
//		n := a[i]
//		x := int64(n)
//		b = dectohex(x)
//		c += strings.Join([]string{b}, "")
//	}
//	fmt.Println(c)
//}
//
//
//
//func DecHex(dec []uint8) string {
//	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70} //声明并初始化一个key和value都是int64的map
//	s := ""
//	for i :=0; i<3; i++ {
//		n := dec[i]
//		x := int64(n)
//		fmt.Println(x)
//		for q := x; q > 0; q = q / 16 {
//			m := q % 16
//			if m > 9 && m < 16 {
//				m = hex[m]
//				s = fmt.Sprintf("%v%v", string(m), s)
//				continue
//			}
//			s = fmt.Sprintf("%v%v", m, s)
//		}
//	}
//	fmt.Println(s)
//	if len(s) == 1{
//		s = "0" + s
//	}
//	return s
//}
//
//func dectohex(n int64) string {
//	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
//	s := ""
//	for q := n; q > 0; q = q / 16 {
//		m := q % 16
//		if m > 9 && m < 16 {
//			m = hex[m]
//			fmt.Println(m)
//			s = fmt.Sprintf("%v%v", string(m), s)
//			continue
//		}
//		fmt.Println(m)
//		s = fmt.Sprintf("%v%v", m, s)
//	}
//	if len(s) == 1{
//		s = "0" + s
//	}
//	return s
//}
var n = "123";

func A(m *string){
	fmt.Println(m)
	n = "456"
	fmt.Println(n)
}

func main()  {
	A(&n)
	fmt.Println(n)
}



