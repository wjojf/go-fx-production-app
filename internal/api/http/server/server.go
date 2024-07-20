package server

import (
	"github.com/gofiber/fiber/v3"
)

func New(options ...Option) *fiber.App {
	// Create a new Fiber instance
	app := fiber.New()

	// Register middleware
	for _, option := range options {
		app = option(app)
	}

	return app
}
