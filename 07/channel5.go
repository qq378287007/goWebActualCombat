package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "i"
	ch <- "love"
	ch <- "you"
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
