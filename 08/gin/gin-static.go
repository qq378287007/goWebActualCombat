package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/assets", "/var/www/gin/assets")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 启动服务
	router.Run(":8080")
}
