package main

import (
	"fmt"
	"os"
)

func main() {
	defer os.Remove("os_write_to_file.txt")

	//新建文件
	fout, err := os.Create("os_write_to_file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\r\n", "Hello Go", i) //Sprintf控制台输出，并有返回值string
		// 写入文件
		fout.WriteString(outstr)            //string信息
		fout.Write([]byte("i love go\r\n")) //byte类型
	}
}
