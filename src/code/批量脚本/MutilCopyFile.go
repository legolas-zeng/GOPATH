package main

import "fmt"
import "math/rand"

const (
	SoureData = "C:\\Users\\Administrator.000\\Desktop\\xiaojiejie\\SoureData"
	OutputData = "C:\\Users\\Administrator.000\\Desktop\\xiaojiejie\\OutputData"
)


func getRand(ch chan int) {
	ch <- rand.Intn(5)
}
func printRand(ch chan int) {
	fmt.Println("Rand Number = ", <-ch)
}
func main() {
	rand.Seed(63)
	ch := make(chan int)
	for i := 1; i <= 5; i++ {
		go getRand(ch)
		fmt.Println(i)
		go printRand(ch)
	}
}