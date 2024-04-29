package main

import (
	"fmt"

	"./person"
)

func main() {
	s := new(person.Student)
	s.name = "shirdon"
	s.Age = 22
	fmt.Println(s.Age)
}
