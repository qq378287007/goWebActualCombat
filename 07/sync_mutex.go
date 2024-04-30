package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()
	for i := 1; i <= 5; i++ {
		wait.Add(1)
		//go func(id int) {
		//	fmt.Println("Not lock:", i)
		//	mutex.Lock()
		//	fmt.Println("Lock:", i)
		//	time.Sleep(time.Second)
		//	fmt.Println("Unlock:", i)
		//	mutex.Unlock()
		//	defer wait.Done()
		//}(i)
		var f = func(id int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock:", i)
			mutex.Unlock()
			defer wait.Done()
		}
		go f(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait()
}
