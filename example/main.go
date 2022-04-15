package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	hostname, _ := os.Hostname()

	router := app.Group("/example")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from example service: " + hostname)
	})

	app.Listen(":80")
}
