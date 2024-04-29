package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		c.SetCookie("my_cookie", "cookievalue", 3600, "/", "localhost", false, true) // 设置cookie
	})

	r.Run(":8068")
}

func Handler(c *gin.Context) {
	// 根据cookie名字读取cookie值
	data, err := c.Cookie("my_cookie")
	if err != nil {
		// 直接返回cookie值
		c.String(200, data)
		return
	}
	c.String(200, "not found!")
}
