package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hi", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Header("site", "shirdon")
	})

	r.Run(":8068")
}
