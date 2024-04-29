package main

import (
	"fmt"
)

func Counting(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	channels := make([]chan int, 6)
	for i := 0; i < 6; i++ {
		channels[i] = make(chan int)
		go Counting(channels[i])
	}

	for _, ch := range channels {
		<-ch
	}

	//select {
	//case <- chan1:
	//	// 如果chan1成功读到数据
	//
	//case chan2 <- 1:
	//	// 如果成功向chan2写入数据
	//
	//default:
	//	// 默认分支
	//}
	//
	//timeout := make(chan bool, 1)
	//
	//go func() {
	//	time.Sleep(1e9)
	//	timeout <- true
	//}()
	//
	//switch {
	//case <- ch:
	//	// 从ch中读取到数据
	//
	//case <- timeout:
	//	// 没有从ch中读取到数据，但从timeout中读取到了数据
	//}
}
