package controllers

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"

	"fmt"
	"context"
	_"time"
	"github.com/astaxie/beego"
	"time"
)


// 创建远程连接的客户端
func newClient() (*client.Client, error) {
	// cli, err := client.NewClientWithOpts()
	hosts := "tcp://"+":2375"     //拼接字符串
	fmt.Println(hosts)
	cli, err := client.NewClient("tcp://192.168.3.5:2375", "1.24", nil, nil)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return cli, err
}

// 列出镜像
func listImages(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	beego.Info(err)

	for _, image := range images {
		fmt.Println(image)
	}
}

// 运行容器
func StartContainer(containerId string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	return err
}

// 停止容器
func StopContainer(containerID string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	timeout := time.Second * 5
	err = cli.ContainerStop(context.Background(), containerID, &timeout)
	return err
}

// 重启容器
func ReStartContainer(containerID string) error {
	cli, err := newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	timeout := time.Second * 5
	err = cli.ContainerRestart(ctx,containerID,&timeout)
	return err
}
//检查容器
func CheckContainer(containerID string) (error) {
	cli,err :=newClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	//container,err := cli.ContainerInspect(ctx,containerID)
	container,err :=cli.ContainerStats(ctx,containerID,true)
	fmt.Println(container)
	return err
}
//// 停止
//func StopContainer(containerID string, cli *client.Client) string {
//	timeout := time.Second * 10
//	err := cli.ContainerStop(context.Background(), containerID, &timeout)
//	if err != nil {
//		beego.Info(err)
//		status := "1"
//		return status
//	} else {
//		fmt.Printf("容器%s已经被停止\n", containerID)
//		status := "0"
//		return status
//	}
//}

// 显示容器列表
func ListContainers() ([]types.Container, error) {
	cli, err := newClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	return containers, nil
}




