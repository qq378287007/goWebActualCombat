package main

import (
	"os"
)

func main() {
	defer os.Remove("write3.txt")

	file, err := os.Create("write3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Go Web编程实战派从入门到精通")
}
