package main

import (
	"fmt"
)

func main() {
	printString()
	fmt.Println("I love go build!")
}

// go mod init main & go mod tidy
// go build & main.exe

// go build main.go utils.go & main.exe
