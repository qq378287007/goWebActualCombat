package main

import (
	"fmt"
	"os"
)

func main() {
	defer os.Remove("open2.txt")

	// 读写方式打开文件
	fp, err := os.OpenFile("open2.txt", os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("文件打开失败。")
		return
	}

	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
}
