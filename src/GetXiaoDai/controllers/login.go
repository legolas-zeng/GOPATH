package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
    "GetXiaoDai/models"
    "strings"
    "github.com/astaxie/beego/orm"
)

type LoginController struct{
    beego.Controller
}

func (this *LoginController) Index(){
    //获取session值
    this.Data["username"] = this.GetSession("username")
    this.TplName = "login.html"
    this.Render()
}

// 登录页面 post
func (this *LoginController) HandleLogin() {
   // 拿到浏览器数据，并去除两边的空格
   Name := strings.TrimSpace(this.GetString("userName"))
   Pwd := strings.TrimSpace(this.GetString("passWord"))
   beego.Info("账号:", Name, "密码:", Pwd)
   if Name == "" || Pwd == "" {
       this.TplName = "login.html"
       this.Data["errmsg"] = "登录失败！！！！！"
       this.Render()
   }
   o := orm.NewOrm()
   var user models.User
   user.Name = Name
   err := o.Read(&user, "name")
   beego.Info(user)
   if err != nil {
       this.Data["errmsg"] = "用户名不存在！！！！！"
       this.TplName="login.html"
       this.Render()
   }else {
       if user.Pwd != Pwd {
           fmt.Println("密码登录失败！！！")
           this.Data["errmsg"] = "密码登录失败！！"
           this.TplName="login.html"
           this.Render()
       }else {
           this.SetSession("username",user.Name)
           this.Redirect("/", 302)
       }
   }
}

func (this *LoginController) Register(){
    u := models.User{}
    //处理表单提交的数据
    if err := this.ParseForm(&u); err != nil{
        fmt.Println(err)
    } else {
        //注册session值
        this.SetSession("uid", u.Name)
    }
    this.Data["uid"] = this.GetSession("uid")
    this.TplName = "login.html"
    this.Render()
}

func (this *LoginController)HandleLogout(){
    this.DelSession("username")
    this.Redirect("/",302)

}