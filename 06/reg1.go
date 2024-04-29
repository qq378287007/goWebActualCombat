package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`Go(\w+)`)
	fmt.Println(re.ReplaceAllString("Hello Gopher，Hello GoLang", "Java$1"))
}
