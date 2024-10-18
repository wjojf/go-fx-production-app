package service

import "go.uber.org/fx"

var Module = fx.Module(
	"http-services",

	fx.Provide(NewJwtService),
)
