package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaffarell/golang-microservice-template/lib"
)

// ExampleController data type
type ExampleController struct {
	logger lib.Logger
}

// NewExampleController creates new user controller
func NewExampleController(logger lib.Logger) ExampleController {
	return ExampleController{
		logger: logger,
	}
}

// GetUser gets the user
func (u ExampleController) Hello(c *gin.Context) {
	u.logger.Info("Hello World from controller")
	c.JSON(200, gin.H{"data": "Hello"})
}
