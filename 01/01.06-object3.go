package main

import (
	"fmt"

	"01/person"
)

func main() {
	s := new(person.Student)
	//s.name = "shirdon"
	s.Age = 22
	fmt.Println(s.Age)
}
