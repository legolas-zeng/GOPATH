package main

import (
    "fmt"
    "github.com/go-redis/redis"
)


func main(){
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379", // redis地址
        Password: "", // redis密码，没有则留空
        DB:       0,  // 默认数据库，默认是0
    })
    sub := client.Subscribe("result")
    fmt.Println(sub.Receive())
    for msg := range sub.Channel(){
        fmt.Println(msg.Payload)

    }
}
