package main

import "fmt"

func main(){
    b :=1111
    c :=&b  //获取b的地址c的类型时*int
    test(c)
    fmt.Println(b)  //值为333发送了变化
}

func test(a *int){
    *a=333
}
////可以与下面进行对比
//func main(){
//    b :=1111
//    test(b)
//    fmt.Println(b)
//}
//func test(a int){  //如果不是传入地址,他就会开辟一个新的内存空间对于原来值没有影响
//    a=333
//}
