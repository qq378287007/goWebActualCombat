package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// 返回 JSON 格式数据
	greeting := Greeting{
		"欢迎一起学习《Go Web编程实战派从入门到精通》",
	}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
