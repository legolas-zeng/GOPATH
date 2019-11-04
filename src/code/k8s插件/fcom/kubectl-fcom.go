package main

import (
	"github.com/mritd/promptx"
	"os/exec"
	"fmt"
)

type CommandType string

const (
	GETPOD CommandType = "getpod"
	GETALL CommandType = "getall"
	GETPV  CommandType = "getpv"
	GETPVC  CommandType = "getpvc"
)

type TypeCommand struct {
	CommandName   CommandType
	ZHDescription string
	K8sCommand    string
}

func main() {
	CommandMsg := []TypeCommand{
		{CommandName: GETPOD, ZHDescription: "获取pod详情", K8sCommand: "kubectl get pod -o wide"},
		{CommandName: GETALL, ZHDescription: "获取all详情", K8sCommand: "kubectl get pod -o wide --all-namespaces"},
		{CommandName: GETPV, ZHDescription: "获取pv详情", K8sCommand: "kubectl get pv"},
		{CommandName: GETPVC, ZHDescription: "获取pvc详情", K8sCommand: "kubectl get pvc"},
	}
	cfg := &promptx.SelectConfig{
		ActiveTpl:    "» {{ .CommandName | cyan }} ({{ .K8sCommand | cyan }})",
		InactiveTpl:  "  {{ .CommandName | white }} ({{ .K8sCommand | white }})",
		SelectPrompt: ">>>>>>>>请选择命令<<<<<<<<",
		SelectedTpl:  "{{ \"» CommandName:\" | green }} {{ .CommandName }}",
		DisPlaySize:  9,
		DetailsTpl: `
--------- 命令详情 ----------
{{ "命令:" | faint }}	{{ .CommandName }}
{{ "详情:" | faint }}	{{ .ZHDescription }}({{ .K8sCommand }})`,
	}

	s := &promptx.Select{
		Items:  CommandMsg,
		Config: cfg,
	}
	//fmt.Println(s.Run())

	selectCommand := s.Run()
	// 打印出k8s命令
	//fmt.Println(CommandMsg[selectCommand].K8sCommand)
	out,err:= exec_shell(CommandMsg[selectCommand].K8sCommand)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

}

func exec_shell(com string) (string ,error){
	command := com
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes,err := cmd.Output()
	resp := string(bytes)
	return resp,err
}

