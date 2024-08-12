package users

import (
	"github.com/wjojf/go-uber-fx/internal/domain/users/events/handlers/verify"
	"go.uber.org/fx"
)

var EventHandlers = fx.Module(
	"user-event-handlers",

	fx.Provide(
		verify.NewHandler,
	),
)
