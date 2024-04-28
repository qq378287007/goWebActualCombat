package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/v1", version1)
	// 路由注册完，开始运行
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// version1 handler
func version1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is a version 1 response!")
}
