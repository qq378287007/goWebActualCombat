package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}

	router.Run(":8080")
}

func login(c *gin.Context) {
}
func submit(c *gin.Context) {
}
func read(c *gin.Context) {
}
