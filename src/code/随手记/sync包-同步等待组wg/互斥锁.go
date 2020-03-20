package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var socket = 10
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
    // a := 1
    // go func() {
    // 	a = 2
    // 	fmt.Println("go a = ", a)
    // }()
    // a = 3
    // time.Sleep(1 * time.Second)
    // fmt.Println("main a = ", a)
    wg.Add(4)
    go sellSocket("售票口1")
    go sellSocket("售票口2")
    go sellSocket("售票口3")
    go sellSocket("售票口4")
    // time.Sleep(3 * time.Second)
    wg.Wait()
    fmt.Println("程序执行完毕")
}

func sellSocket(name string) {
    //defer 语句wg减1
    defer wg.Done()
    rand.Seed(time.Now().UnixNano())
    for {
        mutex.Lock()
        if socket > 0 {
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
            socket--
            fmt.Printf("%v 售出了一张票，剩余%v张票\n", name, socket)
        } else {
            fmt.Println(name, " 亲抱歉，票售完了")
            mutex.Unlock() //条件不满足也要解锁
            break
        }
        mutex.Unlock()
    }
}