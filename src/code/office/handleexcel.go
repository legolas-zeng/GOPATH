package main

import(
"fmt"
"github.com/tealeg/xlsx"
    "os"
    "log"
    "io"
)

var (
    inFile = "C:\\Users\\Administrator\\Desktop\\20200114.xlsx"
)
func main(){
    // 打开文件
    xlFile, err := xlsx.OpenFile(inFile)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    sheet := xlFile.Sheets[0]
    fmt.Println("工作表名: ", sheet.Name)
    for _, row := range sheet.Rows[1:] {
        number := row.Cells[0]
        filename := row.Cells[7]
        //name := row.Cells[2]
        //path := row.Cells[8]
        fullname := fmt.Sprintf("%s-%s", number, filename)
        fmt.Println(fullname)
        func_log2file(fullname)
        //for _, cell := range row.Cells {
        //    fmt.Println(cell)
        //}
        fmt.Print("\n")
    }
    fmt.Println("\n读取成功")
    //func_log2file()
    //func_log2fileAndStdout()
}

func func_log2file(msg string) {
    //创建日志文件
    f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }
    //完成后，延迟关闭
    defer f.Close()
    // 设置日志输出到文件
    log.SetOutput(f)
    // 写入日志内容
    log.Println("check to make sure it works")
}
func func_log2fileAndStdout(msg string) {
    //创建日志文件
    f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }
    //完成后，延迟关闭
    defer f.Close()
    // 设置日志输出到文件
    // 定义多个写入器
    writers := []io.Writer{
        f,
        os.Stdout}
    fileAndStdoutWriter := io.MultiWriter(writers...)
    // 创建新的log对象
    logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
    // 使用新的log对象，写入日志内容
    logger.Println(msg)
}