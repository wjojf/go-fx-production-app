package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gql "github.com/wjojf/go-uber-fx/internal/api/graphql/resolvers"
)

func NewServer(r *gql.Resolver) *http.ServeMux {

	mux := http.NewServeMux()

	httpHandler := handler.NewDefaultServer(
		gql.NewExecutableSchema(
			gql.Config{
				Resolvers: r,
			},
		),
	)

	mux.Handle("POST /graphql", httpHandler)
	mux.Handle("GET /graphql", playground.Handler("GraphQL playground", "/graphql"))

	return mux
}
