package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "sync"
)

func NewRedisClient() (conn redis.Conn, err error) {
    host := "127.0.0.1"
    port := "6379"
    adderss := host + ":" + port
    c, err := redis.Dial("tcp", adderss)
    return c, err
}

func Publish()  {
    conn, err := NewRedisClient()
    if err != nil {
        return
    }
    type Data struct {
        Name *string
        Age *int
    }
    _, err = conn.Do("Publish", "order-create", "1111111111111")
    if err != nil {
        fmt.Println("发布错误", err)
        return
    }
}

func main()  {
    var wg sync.WaitGroup
    wg.Add(1)
    Publish()
    wg.Wait()
}

