package main

import (
	"log"
	"net/http"

	"gitee.com/shirdonl/goWebActualCombat/chapter3/controller"
)

func main() {
	http.HandleFunc("/getUser", controller.UserController{}.GetUser)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
