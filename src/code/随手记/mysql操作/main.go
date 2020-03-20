package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "time"
)

const (
    USERNAME = "root"
    PASSWORD = "qq1005521"
    NETWORK  = "tcp"
    SERVER   = "localhost"
    PORT     = 3306
    DATABASE = "xiaodai"
)

type User struct {
    ID      int64           `db:"id"`
    Name    sql.NullString  `db:"name"`  //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
    Age     int             `db:"age"`
}

type Xiaodai struct {
    XiaodaiId     int           `db:"id"`
    FileName      string        `db:"file_name"`
    HashName      string        `db:"hash_name"`
    FilePath      string        `db:"file_path"`
    Path          string        `db:"path"`
}

//单行查询
func queryOne(DB *sql.DB){
    xiaodai := new(Xiaodai)
    row := DB.QueryRow("select * from xiao_dai where id=?",1)
    //row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
    if err :=row.Scan(&xiaodai.XiaodaiId,&xiaodai.FileName,&xiaodai.HashName,&xiaodai.FilePath,&xiaodai.Path); err != nil{
        fmt.Printf("scan failed, err:%v",err)
        return
    }
    fmt.Println(*xiaodai)
}

//查询多行
func queryMulti(DB *sql.DB){
    user := new(User)
    rows, err := DB.Query("select * from xiao_dai where id > ?", 1)
    defer func() {
        if rows != nil {
            rows.Close() //可以关闭掉未scan连接一直占用
        }
    }()
    if err != nil {
        fmt.Printf("Query failed,err:%v", err)
        return
    }
    for rows.Next() {
        err = rows.Scan(&user.ID, &user.Name, &user.Age) //不scan会导致连接不释放
        if err != nil {
            fmt.Printf("Scan failed,err:%v", err)
            return
        }
        fmt.Print(*user)
    }
}

func insertData(DB *sql.DB){
    result,err := DB.Exec("insert INTO xiao_dai(id,file_name,hash_name,file_path,path) values(?,?,?,?,?)",15,"YDZ","abcnd","podd","odsd")
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    lastInsertID,err := result.LastInsertId()  //插入数据的主键id
    if err != nil {
        fmt.Printf("Get lastInsertID failed,err:%v",err)
        return
    }
    fmt.Println("LastInsertID:",lastInsertID)
    rowsaffected,err := result.RowsAffected()  //影响行数
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

func updateData(DB *sql.DB){
    result,err := DB.Exec("UPDATE users set age=? where id=?","30",3)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

func deleteData(DB *sql.DB){
    result,err := DB.Exec("delete from users where id=?",1)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
    lastInsertID,err := result.LastInsertId()
    if err != nil {
        fmt.Printf("Get lastInsertID failed,err:%v",err)
        return
    }
    fmt.Println("LastInsertID:",lastInsertID)
    rowsaffected,err := result.RowsAffected()
    if err != nil {
        fmt.Printf("Get RowsAffected failed,err:%v",err)
        return
    }
    fmt.Println("RowsAffected:",rowsaffected)
}

func main() {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
    DB, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("Open mysql failed,err:%v\n", err)
        return
    }
    DB.SetConnMaxLifetime(100 * time.Second)
    DB.SetMaxOpenConns(100)
    DB.SetMaxIdleConns(16)
    //queryOne(DB)
    //queryMulti(DB)
    insertData(DB)
    //updateData(DB)
    //deleteData(DB)
}
