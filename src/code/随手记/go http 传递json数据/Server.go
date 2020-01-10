package main

import (
    "encoding/json"
    "fmt"
    "html"
    "io/ioutil"
    "log"
    "net/http"

)


type Cmd struct {
    ReqType int
    FileName string

}

func main() {

    http.HandleFunc("/bar", func (w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

        if r.Method == "POST" {
            b, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Println("Read failed:", err)
            }
            defer r.Body.Close()

            cmd := &Cmd{}
            err = json.Unmarshal(b, cmd)
            if err != nil {
                log.Println("json format error:", err)
            }

            log.Println("cmd:", cmd)
        } else {

            log.Println("ONly support Post")
            fmt.Fprintf(w, "Only support post")
        }

    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}