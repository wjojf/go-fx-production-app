package logging

import "go.uber.org/fx"


var Module = fx.Module(
	"logging",
	fx.Provide(SetupLogger),
)
