package main

import(
"fmt"
"github.com/tealeg/xlsx"
)

var (
    inFile = "C:\\Users\\Administrator.000\\Desktop\\20200114.xlsx"
)
func main(){
    // 打开文件
    xlFile, err := xlsx.OpenFile(inFile)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    // 遍历sheet页读取
    //for _, sheet := range xlFile.Sheets {
    //    fmt.Println("sheet name: ", sheet.Name)
    //    //遍历行读取
    //    for _, row := range sheet.Rows {
    //        // 遍历每行的列读取
    //        for _, cell := range row.Cells {
    //            text := cell.String()
    //            fmt.Printf("%20s", text)
    //        }
    //        fmt.Print("\n")
    //    }
    //}
    sheet := xlFile.Sheets[0]
    fmt.Println("工作表名: ", sheet.Name)
    for _, row := range sheet.Rows[1:] {
        number := row.Cells[0]
        filename := row.Cells[7]
        //name := row.Cells[2]
        //path := row.Cells[8]
        fullname := fmt.Sprintf("%s-%s", number, filename)
        fmt.Println(fullname)
        //for _, cell := range row.Cells {
        //    fmt.Println(cell)
        //}
        fmt.Print("\n")
    }
    fmt.Println("\n读取成功")
}