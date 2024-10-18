package auth

import "go.uber.org/fx"

var Routes = fx.Module(
	"auth-routes",
	fx.Invoke(AddRoutes),
)
