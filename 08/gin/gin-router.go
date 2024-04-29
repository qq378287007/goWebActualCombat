package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 创建v1组
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
	}
	// 创建v2组
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
	}
	router.Run(":8080")
}
