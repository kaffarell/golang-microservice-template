package lib

import (
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type DefaultGinEngine struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewDefaultGinEngine(logger Logger) DefaultGinEngine {
	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.Default()
	return DefaultGinEngine{Gin: engine}
}
