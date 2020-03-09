package main

import (
    "github.com/gomodule/redigo/redis"
    "sync"
    "fmt"
)

func newRedisclient() (conn redis.Conn, err error) {
    host := "127.0.0.1"
    port := "6379"
    adderss := host + ":" + port
    c, err := redis.Dial("tcp", adderss)
    return c, err
}

func resolveOrderCreate(wait *sync.WaitGroup)  {
    defer wait.Done()
    conn, err := newRedisclient()
    if err != nil {
        return
    }
    client := redis.PubSubConn{conn}
    err = client.Subscribe("order-create")
    if err != nil {
        fmt.Println("订阅错误:", err)
        return
    }
    fmt.Println("等待订阅数据 ---->")
    for {
        switch v := client.Receive().(type){
        case redis.Message:
            fmt.Println("Message", v.Channel, string(v.Data))
        case redis.Subscription:
            fmt.Println("Subscription", v.Channel, v.Kind, v.Count)
        }
    }
}

func main()  {
    var wg sync.WaitGroup
    wg.Add(1)
    go resolveOrderCreate(&wg)
    wg.Wait()
}