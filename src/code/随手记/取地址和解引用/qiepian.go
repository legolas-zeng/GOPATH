package main

import "fmt"

func main(){
   b :=&[]int{1,2,3}
   Test(b)
   fmt.Println(b)
   c :=[]int{4,5,6}
   Test2(c)
    fmt.Println(c)
}
func Test(a *[]int){
   (*a)[1]=3
}

//如果传入对象是值类型,不是引用类型这个不生效,只正对引用类型切片才生效,数组值类型不生效,只能按照方式一写
func Test2(a []int){
    a[1]=6
}

//GO对于切片做了优化可以省略写内容


//func main(){
//  b :=[]int{1,2,3}
//  Test(b)
//  fmt.Println(b)
//}
//func Test(a []int){
//  a[1]=3
//}

