package main

import "fmt"

func main() {

	var a, b int
	go func() {
		a = 1
		fmt.Println("a:", b, " ")
	}()
	go func() {
		b = 1
		fmt.Println("b:", a, " ")
	}()
}
