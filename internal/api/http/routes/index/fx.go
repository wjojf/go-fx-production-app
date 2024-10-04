package index

import "go.uber.org/fx"

var Route = fx.Module(
	"index-route",
	fx.Invoke(AddRoute),
)
