package main

import (
	"fmt"
	"os"
)

func main() {
	defer os.Remove("write2.txt")

	file, err := os.Create("write2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Go Web编程实战派从入门到精通")
	n, err := file.WriteAt([]byte("Go语言Web"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
