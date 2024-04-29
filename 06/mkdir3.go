// ++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer os.RemoveAll("static")

	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		fmt.Println(err)
	}
}
