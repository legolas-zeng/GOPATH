package main

import "fmt"

// 接口可以理解为一些行为、方法的集合。
// 定义 TypeCalculator 为任何具有 TypeCal 方法的类型。TypeCal 方法没有参数，返回字符串,所有定义了该方法的类型我们称它实现了 TypeCalculator 接口
type TypeCalculator interface {
    // 定义一个待实现的方法TypeCal()
    TypeCal() string
}

type Worker struct {
    Type int
    Name string
}

type Student struct {
    Name string
}

func (w Worker) TypeCal() string {
    if w.Type == 0 {
        return w.Name +"是蓝翔毕业的员工"
    } else {
        return w.Name+"不是蓝翔毕业的员工"
    }
}

func (s Student) TypeCal() string  {
    return s.Name + "还在蓝翔学挖掘机炒菜"
}

func main() {
    // 创建两个struct结构体的实例
    worker := Worker{Type:0, Name:"小华"}
    student := Student{Name:"小明"}
    // 把这两个实例放同一个TypeCalculator的切片中
    workers := []TypeCalculator{worker, student}
    // 遍历这个切片，并调用切片中的函数打印结果
    for _, v := range workers {
        fmt.Println(v.TypeCal())
    }
}