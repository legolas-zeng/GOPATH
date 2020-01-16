package main

import "fmt"

func main() {
    var out []*int
    for i := 0; i < 3; i++ {
        out = append(out, &i)
    }
    fmt.Println("Values:", *out[0], *out[1], *out[2])
    fmt.Println("Addresses:", out[0], out[1], out[2])

    var out2 []*int
    for i := 0; i < 3; i++ {
        i := i // Copy i into a new variable.
        out2 = append(out2, &i)
    }
    fmt.Println("Values:", *out2[0], *out2[1], *out2[2])
    fmt.Println("Addresses:", out2[0], out2[1], out2[2])
}

//因为每次循环中，我们只是把变量 i 的地址放进 out 数组里，因为变量 i 是同一个变量，只有在循环结束的时候，被赋值为3。
//
//解决方法：申明一个新的变量


