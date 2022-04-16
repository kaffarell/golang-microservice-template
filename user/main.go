package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user_service := router.Group("/user")

	hostname, _ := os.Hostname()

	user_service.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "User service: "+hostname)
	})

	router.Run(":80")
}
