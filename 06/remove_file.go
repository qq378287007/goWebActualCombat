package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为“test_rename”的目录，perm权限为0777
	err := os.Mkdir("test_remove", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("created dir:test_remove")

	//创建一个名为“test_remove1.txt”的空文件
	fp1, err := os.Create("./test_remove/test_remove1.txt") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println(err)
	}
	fp1.Close()
	fmt.Println("created file:test_remove1.txt")

	fp2, err := os.Create("./test_remove/test_remove2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fp2.Close()
	fmt.Println("created file:test_remove2.txt")

	fp3, err := os.Create("./test_remove/test_remove3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fp3.Close()
	fmt.Println("created file:test_remove3.txt")

	err = os.Remove("./test_remove/test_remove1.txt")
	if err != nil {
		fmt.Printf("removed ./test_remove/test_remove1.txt err : %v\n", err)
	}
	fmt.Println("removed file:./test_remove/test_remove1.txt")

	err = os.RemoveAll("./test_remove")
	if err != nil {
		fmt.Printf("remove all ./test_remove err : %v\n", err)
	}
	fmt.Println("removed all files:./test_remove")
}
