package main

import (
	"fmt"
	"os"
)

func main() {
	defer os.Remove("chmod.txt")

	fp, err := os.Create("chmod.txt") // 创建文件, 默认权限0666-rw-r--r--, 存在会清空。
	if err != nil {
		fmt.Println("文件创建失败。")
		return
	}
	defer fp.Close() //关闭文件，释放资源。

	fileInfo, err := os.Stat("chmod.txt")
	fileMode := fileInfo.Mode()
	fmt.Println(fileMode)
	os.Chmod("chmod1.txt", 0777) //重新赋权限-rwxrwxrwx
	fileInfo, err = os.Stat("chmod.txt")
	fileMode = fileInfo.Mode()
	fmt.Println(fileMode)
}
