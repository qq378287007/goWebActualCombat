package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("./move_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fp.Close()

	os.Mkdir("tmp", 0777)
	err = os.Rename("./move_file.txt", "./tmp/move_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	os.RemoveAll("tmp")
}
