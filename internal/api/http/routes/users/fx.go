package users

import "go.uber.org/fx"

var Routes = fx.Module(
	"users-routes",
	fx.Invoke(
		AddRoutes,
	),
)
