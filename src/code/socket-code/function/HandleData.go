package function

import (
    "fmt"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "strings"
)

const (
    USERNAME = "zwa"
    PASSWORD = "qq1005521"
    NETWORK  = "tcp"
    SERVER   = "192.168.3.5"
    PORT     = 3306
    DATABASE = "pcinfo"
)

type Pcinfo struct {
    Id          int                  `db:"id"`
    Ip          sql.NullString       `db:"ip"`
    Cpu         sql.NullString       `db:"cpu"`
    OsInfo      sql.NullString       `db:"osinfo"`
    Men         sql.NullString       `db:"men"`
    Online      int                  `db:"online"`
    Port        sql.NullString       `db:"port"`
    Pcname      sql.NullString       `db:"pcname"`
}

// mysql数据库初始化
func mysqlConn() *sql.DB{
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
    DB, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("Open mysql failed,err:%v\n", err)
    }
    DB.SetConnMaxLifetime(100 * time.Second)
    DB.SetMaxOpenConns(100)
    DB.SetMaxIdleConns(16)
    return DB
}

//插入电脑硬件数据
func UpdatePcData(pcinfo map[string]string){
    DB := mysqlConn()
    _,err := DB.Exec("UPDATE pcinfo set cpu=?,osinfo=?,men=?,pcname=? where ip=?",pcinfo["cpu"],pcinfo["osinfo"],pcinfo["men"],pcinfo["pcname"],pcinfo["ip"])
    if err != nil{
        fmt.Printf("更新失败,err:%v",err)
        return
    }
    DB.Close()
}

//初始化客户端，判断是否存在记录
func InitPcInfo(clientip string){
    add := strings.Split(clientip,":")
    req:=queryOne(add[0])
    DB := mysqlConn()
    fmt.Println(req)
    if req ==0 {
    _,err := DB.Exec("insert INTO pcinfo(ip,port,online) values(?,?,?)",add[0],add[1],0)
    //先判断是否存在这个ip的数据
    if err != nil{
       fmt.Printf("初始化数据失败,err:%v",err)
       return
    }
    }else if req==1{
        _,err := DB.Exec("UPDATE pcinfo set online=?,port=? where ip=?",0,add[1],add[0])
        if err != nil{
            fmt.Printf("更新状态失败,err:%v",err)
            return
        }
    }
    DB.Close()
}

func OfflineClient(clientip string)  {
    DB := mysqlConn()
    _,err := DB.Exec("UPDATE pcinfo set online=? where ip=?",1,clientip)
    if err != nil{
        fmt.Printf("Insert failed,err:%v",err)
        return
    }
}

func queryOne(ip string) int{
    DB := mysqlConn()
    rows:= DB.QueryRow("select * from pcinfo where ip=?",ip)
    pcinfo := new(Pcinfo)
    err := rows.Scan(&pcinfo.Id, &pcinfo.Ip, &pcinfo.Cpu, &pcinfo.OsInfo, &pcinfo.Men, &pcinfo.Online, &pcinfo.Port,&pcinfo.Pcname)
    if err == nil {
        fmt.Println(pcinfo.Id, pcinfo.Ip, pcinfo.Cpu, pcinfo.OsInfo, pcinfo.Men, pcinfo.Online, pcinfo.Port,pcinfo.Pcname)
        return 1
    }else {
        fmt.Println("没有结果")
        return 0
    }
}
