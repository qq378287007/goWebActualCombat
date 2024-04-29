package main

import (
	"fmt"
	"runtime"
	"sync"
)

var Counter int

func Count(lock *sync.Mutex) {
	lock.Lock() // 上锁
	Counter++
	fmt.Println("Counter =", Counter)
	lock.Unlock() // 解锁
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 6; i++ {
		go Count(lock)
	}
	for {
		lock.Lock() // 上锁
		c := Counter
		lock.Unlock()     // 解锁
		runtime.Gosched() // 出让时间片
		if c >= 6 {
			break
		}
	}
}
