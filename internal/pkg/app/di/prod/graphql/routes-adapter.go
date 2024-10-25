package graphql

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/api/graphql/adapters"
)

func AddGraphQLRoutes(
	app *fiber.App,
	adapter adapters.FiberGraphQLAdapter,
	playgroundAdapter adapters.FiberGraphQLPlaygroundAdapter,
) {
	app.Post("/graphql", adapter)
	app.Get("/graphql", playgroundAdapter)
}
