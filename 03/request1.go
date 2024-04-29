package main

import "net/url"

func main() {
	path := "http://lcoalhost:8082/article?id=1"
	p, _ := url.Parse(path) // 解析url
	println(p.Host)
	println(p.User)
	println(p.RawQuery)
	println(p.RequestURI())
}
