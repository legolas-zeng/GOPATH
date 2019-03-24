package main

import "fmt"
import "math/rand"

const (
	SoureData = "C:\\Users\\Administrator.000\\Desktop\\xiaojiejie\\SoureData"
	OutputData = "C:\\Users\\Administrator.000\\Desktop\\xiaojiejie\\OutputData"
)


//func getRand(ch chan int) {
//	ch <- rand.Intn(5) //把rand值存入ch中
//}
//func printRand(ch chan int) {
//	fmt.Println("Rand Number = ", <-ch) //从ch中读出一个值
//}
//func main() {
//	rand.Seed(63) //不同的随机算子，才能保证随机数是随机的。
//	ch := make(chan int)
//	for i := 1; i <= 5; i++ {
//		go getRand(ch)
//		go printRand(ch)
//	}
//}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case ch <- rand.Intn(5): // 把rand值存入ch中
			case <-done: // If receive signal on done channel - Return
				return
			default:
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("随机数字是：", <-ch) // Print number received on standard output
		}
		done <- true // Send Terminate Signal and return
		return
	}()
	<-done // Exit Main when Terminate Signal received
}

