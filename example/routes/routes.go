package routes

import (
	"github.com/kaffarell/golang-microservice-template/lib"
)

type ExampleRoutes struct {
	logger            lib.logger
	handler           lib.RequestHandler
	exampleController controllers.ExampleController
}

func (e ExampleRoutes) Setup() {
	e.logger.Info("Setting up routes")
	api := e.handler.Gin.Group("/")
	api.GET("/", e.exampleController.Hello)
}

func NewExampleRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	exampleController controllers.ExampleController,
) ExampleRoutes {
	return &ExampleRoutes{
		logger:            logger,
		handler:           handler,
		exampleController: exampleController,
	}
}
