package domain

import (
	"github.com/wjojf/go-uber-fx/internal/domain/users"
	"go.uber.org/fx"
)

var eventHandlers = fx.Module(
	"domain-event-handlers",

	users.EventHandlers,
)
