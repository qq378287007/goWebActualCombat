package main

import (
	"log"
	"net"
	"time"
)

func startConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		log.Printf("%d: dial error: %s", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}

func main() {
	var sl []net.Conn
	for i := 1; i < 500; i++ {
		conn := startConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}

	time.Sleep(time.Second * 10000)
}
