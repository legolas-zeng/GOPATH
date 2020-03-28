package main

import (

    "fmt"
    "sync"
)

func consumer(messages <- chan int, shutdown <- chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case message, ok := <- messages:
            //do something.
            if ok {
                fmt.Println(message)
            } else {
                //no data , exit.
                fmt.Println("no data, exit.")
                return
            }
        case _ = <- shutdown:
            //we `re done!
            //shutdown now , messages buffered channel data may be lost.
            fmt.Println("all done!")
            return
        }
    }
}

func main() {
    shutdown := make(chan int)
    messages := make(chan int, 16)

    wg := &sync.WaitGroup{}
    wg.Add(1)
    go consumer(messages, shutdown, wg)
    for i := 0; i < 10; i++ {
        messages <- i
    }
    close(messages)
    fmt.Println("wait!")
    wg.Wait()


}