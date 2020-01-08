package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    log.Print("启动服务...")
    http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprint(w, "Hello! Your request was processed.")
    }, )
    http.HandleFunc("/handle", func(w http.ResponseWriter, _ *http.Request) {
        //client := http.Client{Timeout: 5 * time.Second}
        //resp, error := client.Get("https://www.baidu.com/")
        //if error != nil {
        //    panic(error)
        //}
        fmt.Println("--------",time.Now().Format("2006-01-02 15:04:05"))
        fmt.Fprint(w, time.Now().Format("2006-01-02 15:04:05"))
    })
    log.Print("服务启动完成.....")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

//打开 http://127.0.0.1:8000/home
