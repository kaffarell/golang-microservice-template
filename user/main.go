package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router := app.Group("/user")
	hostname, _ := os.Hostname()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from user service: " + hostname)
	})

	app.Listen(":80")
}
