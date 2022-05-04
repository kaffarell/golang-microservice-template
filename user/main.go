package main

import (
	"net/http"
	"os"

	"github.com/kaffarell/golang-microservices-template/user/logger"
	"github.com/nullseed/logruseq"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	log := logrus.New()
	router.Use(logger.Logger(log), gin.Recovery())

	user_service := router.Group("/user")

	hostname, _ := os.Hostname()

	log.AddHook(logruseq.NewSeqHook("http://seq:5341"))
	log.Info("Application started")

	user_service.GET("/", func(c *gin.Context) {
		log.Info("Cool")
		c.JSON(http.StatusOK, "User service: "+hostname)
	})

	router.Run(":80")
}
