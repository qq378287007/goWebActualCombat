package main

import (
	"fmt"
	"net/http"
)

func Header(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/header", Header)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
