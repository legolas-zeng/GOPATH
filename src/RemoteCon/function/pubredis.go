package function

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "k8s.io/apimachinery/pkg/util/json"
)

func NewRedisClient() (conn redis.Conn, err error) {
    host := "127.0.0.1"
    port := "6379"
    adderss := host + ":" + port
    c, err := redis.Dial("tcp", adderss)
    return c, err
}

func Publish(cmdinfo map[string]string)  error {
    //var wg sync.WaitGroup
    //wg.Add(1)
    conn, _ := NewRedisClient()

    type Data struct {
        Name *string
        Age *int
    }
    //cmdinfo["192.168.10.3"] = "getinfo"
    value,_ := json.Marshal(cmdinfo)
    fmt.Println(value)
    fmt.Printf("%T",value)
    _, err := conn.Do("Publish", "order-create", value)
    //wg.Done()
    //
    //wg.Wait()
    return err
}