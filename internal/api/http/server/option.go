package server

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Option func(app *fiber.App) *fiber.App

func WithMiddleware(middlewares ...fiber.Handler) Option {
	return func(app *fiber.App) *fiber.App {

		for _, mw := range middlewares {
			app.Use(mw)
		}

		return app
	}
}

func WithStatic(path string, root string) Option {
	return func(app *fiber.App) *fiber.App {

		var cfg = static.Config{
			CacheDuration: 1 * time.Hour,
			MaxAge:        3600,
			Browse:        true,
		}

		app.Get(path, static.New(root, cfg))

		return app
	}
}

func WithRequestID() Option {
	return func(app *fiber.App) *fiber.App {
		app.Use(requestid.New())
		return app
	}
}
