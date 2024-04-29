package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用 io/ioutil.ReadFile 方法一次性将文件读取到内存中
	filePath := "read2.go"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		// log.Fatal(err)
		fmt.Printf("读取文件出错：%v", err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
}
