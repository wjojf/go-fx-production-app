package adapters

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	gql "github.com/wjojf/go-uber-fx/internal/api/graphql/resolvers"
)

type FiberGraphQLAdapter fiber.Handler
type FiberGraphQLPlaygroundAdapter fiber.Handler

func Fiber(r *gql.Resolver) FiberGraphQLAdapter {

	httpHandler := handler.NewDefaultServer(
		gql.NewExecutableSchema(
			gql.Config{
				Resolvers: r,
			},
		),
	)

	return adaptor.HTTPHandler(httpHandler)
}

func FiberPlayground() FiberGraphQLPlaygroundAdapter {
	return adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/graphql"))
}
