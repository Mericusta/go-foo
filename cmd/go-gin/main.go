package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.Header("Server", "local")
		c.JSON(200, gin.H{
			"message": "hello, world!",
		})
	})
	r.Run(":8080")
}
