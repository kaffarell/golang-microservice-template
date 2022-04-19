package main

import (
	"context"

	"github.com/kaffarell/golang-microservice-template/example/controllers"
	"github.com/kaffarell/golang-microservice-template/example/routes"
	"github.com/kaffarell/golang-microservice-template/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func run(handler lib.RequestHandler, routes routes.ExampleRoutes) {
	routes.Setup()
	handler.Gin.Run(":80")
}

func main() {
	logger := lib.GetLogger()
	logger.Info("Hello World")
	opts := fx.Options(
		fx.WithLogger(func() fxevent.Logger {
			return logger.GetFxLogger()
		}),
		fx.Invoke(run),
	)
	ctx := context.Background()
	app := fx.New(
		fx.Options(
			fx.Provide(controllers.NewExampleController),
			fx.Provide(routes.NewExampleRoutes),
			lib.Module,
		),
		opts)
	err := app.Start(ctx)

	defer app.Stop(ctx)
	if err != nil {
		logger.Fatal("Critical Error",
			zap.Error(err),
		)
	}
}
