package main

import (
	"fmt"
	"net/http"
)

// 定义一个处理器
type WebHandler struct{}

// 实现Handler接口
func (h *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi, web")
}

func main() {
	handler := WebHandler{}
	//将处理器绑定在服务器上
	server := http.Server{
		Addr:    "127.0.0.1:8083",
		Handler: &handler,
	}
	server.ListenAndServe()
}
