package main

import (
	"log"
	"net/http"
)

func main() {
	// 启动服务器
	srv := &http.Server{Addr: ":8088", Handler: http.HandlerFunc(handle)}
	// 用TLS启动服务器，因为我们运行的是http/2，它必须是与TLS一起运行。
	log.Printf("Serving on https://0.0.0.0:8088")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

// 处理器方法
func handle(w http.ResponseWriter, r *http.Request) {
	// 记录请求协议
	log.Printf("Got connection: %s", r.Proto)
	// 向客户发送一条消息
	w.Write([]byte("Hello this is a HTTP/2 message!"))
}
