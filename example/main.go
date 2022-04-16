package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	hostname, _ := os.Hostname()

	example_service := router.Group("/example")

	example_service.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Example service: "+hostname)
	})

	router.Run(":80")
}
