package controllers

import (
	"github.com/astaxie/beego"
	models2 "rrjc_ops/models"
	"fmt"
	"strings"
)

type DockerListController struct {
	//BaseController    //这个是自己封装的controller
	beego.Controller
}


func (c *DockerListController) DockerList() {
	beego.ReadFromRequest(&c.Controller)

	DockerInfo := &models2.Dockers{}
	docks:= DockerInfo.FindAllDockInfo()


	fmt.Println(docks)
	c.Data["docks"] = docks
	c.TplName = "base/base.html"
	c.Layout = "base/base.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["re_content"] = "docker/docker_list.html"
	c.LayoutSections["js"] = "docker/docker_list_js.html"
	c.Render()
}

func (c *DockerListController) DockerAdd(){
	ip := c.GetString("ip")
	group := c.GetString("group")
	if ip == "" {
		fmt.Printf("ip为空")
	} else if group == "" {
		fmt.Printf("group为空")
	} else {
		groupInfo := &models2.Groups{}
		groupInfo.Ip = ip
		groupInfo.GroupName = group
		err := groupInfo.SaveGroupInfo()
		if err != nil {
			fmt.Printf("保存失败")
		} else {

			fmt.Printf("保存成功")

		}

	}
	c.TplName = "base/base.html"
	c.Layout = "base/base.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["re_content"] = "docker/docker_add.html"
	c.Render()
}

func (this *DockerListController) DockerFunction(){
	m := &models2.Message{}
	m.Msg = "操作成功"
	m.Status = 1

	IpList := this.GetString("ips")
	Active := this.GetString("active")
	//IpHad := this.GetString("ip")
	//fmt.Printf(IpList)
	fmt.Printf(Active)
	//this.Data["data"] = m
	//this.Render()

	//fmt.Println(reflect.TypeOf(IpList).Kind().String())

	ipListStr := strings.Replace(strings.Replace(strings.Replace(IpList,"[","",-1) ,"]","",-1) ,"\"","",-1)
	ipList := strings.Split(ipListStr,",")
	fmt.Println(ipList)

	for _,ip:=range ipList{
		fmt.Println(ip)
		dockerName := &models2.Dockers{} //实例化docker类
		ContainerNameList := dockerName.FindDockerInfo("dockers",ip)
		fmt.Printf("---------------")
		fmt.Printf("%x",ContainerNameList)
		fmt.Println(ContainerNameList)

	}
	if Active == "启动容器" {
		err := StartContainer("61e74732dbe4")
		if err == nil{
			fmt.Printf("容器启动成功")
		} else {
			fmt.Printf("容器启动失败")
		}
	}
	if Active == "关闭容器" {
		err := StopContainer("61e74732dbe4")
		if err == nil{
			fmt.Printf("容器关闭成功")
		} else {
			fmt.Printf("容器关闭失败")
		}
	}
	if Active == "重启容器" {
		err := ReStartContainer("centos7")
		if err == nil{
			fmt.Printf("容器重启成功")
		} else {
			fmt.Printf("容器重启失败")
		}
	}
	if Active == "检查容器" {
		err := CheckContainer("61e74732dbe4")
		if err == nil{
			fmt.Printf("获取容器信息成功")
		} else {
			fmt.Printf("获取容器信息失败")
		}
	}

}