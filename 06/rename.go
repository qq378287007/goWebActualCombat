package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//创建一个名为“dir_name1”的目录，perm权限为0777
	err := os.Mkdir("dir_name1", 0777)
	if err != nil {
		fmt.Println(err)
	}

	oldName := "dir_name1"
	newName := "dir_name2"
	//将dir_name1重命名为dir_name2
	err = os.Rename(oldName, newName)//文件或目录重命名
	if err != nil {
		log.Fatal(err)
	}

	os.Remove(newName)
}
