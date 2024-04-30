package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// 进行读写操作
		// ... ...
	}
}

func main() {
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// 开启协程来处理连接
		go handleConn(c)
	}
}
