package main

import (
	"fmt"
	"regexp"
)

func main() {
	//re := regexp.MustCompile(`[^eovbgramin]`)
	//fmt.Println(re.ReplaceAllStringFunc("I love Go web programming", strings.ToUpper))

	//re := regexp.MustCompile(`w(a*)i`)
	//fmt.Println(re.ReplaceAllLiteralString("-wi-waai-", "T"))
	//fmt.Println(re.ReplaceAllLiteralString("-wi-waai-", "$1"))
	//fmt.Println(re.ReplaceAllLiteralString("-wi-waai-", "${1}"))

	reg := regexp.MustCompile(`^http://blog.sina.com.cn/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := reg.FindStringSubmatch("http://blog.sina.com.cn/2020/08/20/blog_dfcfef2f0102yyd1.html")
	// 返回[]string{}数据类型
	for _, param := range params {
		fmt.Println(param)
	}
}
