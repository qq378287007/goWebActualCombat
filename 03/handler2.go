package main

import (
	"fmt"
	"net/http"
)

func hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/hello1", hello1)         //使用http包默认的多路复用器实例绑定
	http.ListenAndServe("127.0.0.1:8086", nil) //使用http包默认的服务器实例，并开启监听
}
