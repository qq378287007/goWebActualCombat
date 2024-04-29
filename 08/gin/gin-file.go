package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hi", func(c *gin.Context) {
		//通过File函数，直接返回本地文件，参数为本地文件地址。
		//函数说明：c.File("文件路径")
		c.File("/var/www/gin/test.jpg")
	})

	r.Run(":8068")
}
