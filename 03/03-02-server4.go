package main

import (
	"log"
	"net/http"

	"03/controller"
)

func main() {
	http.HandleFunc("/getUser", controller.UserController{}.GetUser)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
