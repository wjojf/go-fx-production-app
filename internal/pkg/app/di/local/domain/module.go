package domain

import "go.uber.org/fx"

var Module = fx.Module(
	"domain",
	eventHandlers,
)
