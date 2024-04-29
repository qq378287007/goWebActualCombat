package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("read1.go")
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
	}
	// 及时关闭文件句柄
	defer file.Close()

	// bufio.NewReader(rd io.Reader) *Reader
	reader := bufio.NewReader(file)//带缓冲方式读取
	// 循环读取文件的内容
	for {
		line, err := reader.ReadString('\n') // 读到一个换行符就结束
		if err == io.EOF {                   // io.EOF表示文件的末尾
			break
		}
		// 输出内容
		fmt.Print(line)
	}
}
