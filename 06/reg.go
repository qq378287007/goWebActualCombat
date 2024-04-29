package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`Go(\w+)`)
	fmt.Println(re.ReplaceAllString("Hello Gopher，Hello GoLang", "Java$1"))

	text := "Hello Gopher，Hello Go Web"
	reg := regexp.MustCompile(`\w+`)
	fmt.Println(reg.MatchString(text))
}
