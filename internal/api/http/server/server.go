package server

import (
	"github.com/gofiber/fiber/v3"
)

func New(middlewares ...fiber.Handler) *fiber.App {
	// Create a new Fiber instance
	app := fiber.New()

	// Register middleware
	for _, mw := range middlewares {
		app.Use(mw)
	}

	return app
}
