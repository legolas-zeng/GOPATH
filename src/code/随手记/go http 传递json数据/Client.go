package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

)

type Cmd struct {

    ReqType int
    FileName string

}


func main() {

    url := "http://127.0.0.1:8080/bar"
    contentType := "application/json;charset=utf-8"

    cmd := Cmd{ReqType: 12, FileName: "plugin"}
    b ,err := json.Marshal(cmd)
    if err != nil {
        log.Println("json format error:", err)
        return
    }

    body := bytes.NewBuffer(b)

    resp, err := http.Post(url, contentType, body)
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


    log.Println("content:", string(content))

}