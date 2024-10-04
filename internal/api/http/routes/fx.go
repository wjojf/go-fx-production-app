package routes

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/routes/index"
	"github.com/wjojf/go-uber-fx/internal/api/http/routes/users"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-routes",
	users.Routes,
	index.Route,
)
