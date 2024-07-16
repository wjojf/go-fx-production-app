package handlers

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers/users"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-handlers",
	fx.Provide(
		users.New,
	),
)
