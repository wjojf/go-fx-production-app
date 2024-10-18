package handlers

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers/auth"
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers/index"
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers/users"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-handlers",

	fx.Provide(
		users.New,
	),

	fx.Provide(
		index.New,
	),

	fx.Provide(
		auth.New,
	),
)
