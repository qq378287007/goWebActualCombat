package main

import (
	"log"
	"os"
)

func main() {
	os.Mkdir("dir1", 0777)
	err := os.Remove("dir1")//只能删除空目录
	if err != nil {
		log.Fatal(err)
	}
}
