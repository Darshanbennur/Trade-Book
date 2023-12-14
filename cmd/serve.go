package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Execute() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})
	r.Run() // listen and serves on localhost:8080
}
