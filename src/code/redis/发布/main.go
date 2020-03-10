package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "sync"
    "k8s.io/apimachinery/pkg/util/json"
)

var(
    cmdinfo = make(map[string]string)
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
    cmdinfo["2"] = "ping 192.168.3.5 -l 32000"
    value,_ := json.Marshal(cmdinfo)
    fmt.Println(value)
    fmt.Printf("%T",value)
    _, err = conn.Do("Publish", "order-create", value)
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

