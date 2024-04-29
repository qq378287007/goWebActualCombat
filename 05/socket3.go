package main

import (
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "http://www.baidu.com:80")
	if err != nil {
		//处理错误
	}
	//读写操作

	conn, err := net.DialTimeout("tcp", ":8085", 2*time.Second)
	if err != nil {
		//处理错误
	}
	//读写操作
}
