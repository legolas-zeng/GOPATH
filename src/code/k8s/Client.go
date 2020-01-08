package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "fmt"
)

type Cmd struct {
    ReqType int
    FileName string

}


func main() {
    url := "http://127.0.0.1:8000/handle"
    http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
        resp, err := http.Get(url)

        if err != nil {
            log.Println("Post failed:", err)
            return
        }

        defer resp.Body.Close()

        content, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Println("Read failed:", err)
            return
        }

        log.Println("返回值:", string(content))
        fmt.Fprint(w, "返回值:", string(content))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))

}