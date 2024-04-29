package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("./file1.txt")
	if err != nil {
		fmt.Printf("remove ./file1.txt err : %v\n", err)
	}

	err = os.RemoveAll("./file2.txt")
	if err != nil {
		fmt.Printf("remove all ./file2.txt err : %v\n", err)
	}
}
