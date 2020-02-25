package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool)
    go func() {
        for {
            select {
            //case的优先级大于default
            // 如果done成功读到数据，则进行该case处理语句
            case <-done:
                fmt.Println("退出携程01")
                return
            default: //default确保发送不被堵塞
                fmt.Println("监控01...")
                time.Sleep(1 * time.Second)
            }
        }
    }()

    go func() {
        for res := range done { //没有消息阻塞状态，chan关闭 for 循环结束
            fmt.Println(res)
        }
        fmt.Println("退出监控03")
    }()

    go func() {
        for {
            select {
            case <-done:
                fmt.Println("退出协程02")
                return
            default:
                fmt.Println("监控02...")
                time.Sleep(1 * time.Second)
            }
        }
    }()
    time.Sleep(3 * time.Second)
    close(done)
    time.Sleep(5 * time.Second)
    fmt.Println("退出程序")
}