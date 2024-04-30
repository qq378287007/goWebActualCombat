package main

import (
	"fmt"
	"net"
)

func main() {
	//创建监听的地址，并且指定udp协议
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr) //创建监听链接
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, raddr, err := conn.ReadFromUDP(buf) //接收客户端发送过来的数据，填充到切片buf中。
	if err != nil {
		return
	}
	fmt.Println("客户端发送：", string(buf[:n]))

	_, err = conn.WriteToUDP([]byte("你好，客户端，我是服务端"), raddr) //向客户端发送数据
	if err != nil {
		fmt.Println("WriteToUDP err:", err)
		return
	}
}
