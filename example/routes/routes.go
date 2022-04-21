package routes

import (
	"github.com/kaffarell/golang-microservice-template/example/controllers"
	"github.com/kaffarell/golang-microservice-template/lib"
)

type ExampleRoutes struct {
	logger            lib.Logger
	engine            lib.DefaultGinEngine
	exampleController controllers.ExampleController
}

func (e ExampleRoutes) Setup() {
	e.logger.Info("Setting up routes")
	// TODO: maybe inject here the service name to use as default route
	api := e.engine.Gin.Group("/example")
	api.GET("/", e.exampleController.Hello)
}

func NewExampleRoutes(
	logger lib.Logger,
	engine lib.DefaultGinEngine,
	exampleController controllers.ExampleController,
) ExampleRoutes {
	return ExampleRoutes{
		logger:            logger,
		engine:            engine,
		exampleController: exampleController,
	}
}
