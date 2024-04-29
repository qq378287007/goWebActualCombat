package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go playTableTennis("Jim", ch)
	go playTableTennis("Jack", ch)
	ch <- 1
	wg.Wait()

}

func playTableTennis(player string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch
		if !ok {
			fmt.Printf("%s has won!\n", player)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed!\n", player)
			close(ch)
			return
		}
		ball++
		fmt.Printf("play %s has hit the ball, his score is %d\n", player, ball)
		ch <- ball
	}
}
