package main

import "github.com/gin-gonic/gin"

func main() {
	// create router
	f := gin.Default()

	// simple test route
	f.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	// start server at localhost:8080
	f.Run(":8080")
}
