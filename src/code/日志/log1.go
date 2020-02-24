package main

import (
    "os"
    "log"
    "fmt"
)

func main() {

    file, err := os.Create("test.log")
    if err != nil {
        log.Fatalln("fail to create test.log file!")
    }
    logger := log.New(file, "", log.Llongfile)


    // 写入文件log格式：/Users/zhou/go/src/zhouTest/test.go:22: 2.Println log with log.LstdFlags ...
    logger.Println("2.Println log with log.LstdFlags ...")

    logger.SetFlags(log.LstdFlags)    // 设置写入文件的log日志的格式

    // 写入log文件格式： 2018/07/31 17:28:21 4.Println log without log.LstdFlags ...
    logger.Println("4.Println log without log.LstdFlags ...")


    fmt.Println("打印")
    logger.Fatal("9.Fatal log without log.LstdFlags ...")
    fmt.Println("Fatal终止了程序，这句不执行！")
}
