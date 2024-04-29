package main

import (
	"fmt"

	"install/pkg"
)

func main() {
	pkg.CallFunc()
	fmt.Println("I love go build!")
}

// go mod init install & go mod tidy

// go env
// 默认安装在%GOPATH%\bin\目录下
// go install
