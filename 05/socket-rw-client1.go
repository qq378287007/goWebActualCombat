package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run socket-rw-client1.go :... content")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8058")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	data := os.Args[1]
	conn.Write([]byte(data))

	time.Sleep(time.Second * 10000)
}
