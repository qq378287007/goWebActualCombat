package main

import "net"

func main() {
	conn, err := net.Dial("tcp", "192.168.0.1:8087")
	conn, err := net.Dial("udp", "192.168.0.2:8088")
	conn, err := net.Dial("ip4:icmp", "www.shirdon.com")
	net.DialTCP()
	println(conn)
	println(err)
}
