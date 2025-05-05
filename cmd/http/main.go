package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default();

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}