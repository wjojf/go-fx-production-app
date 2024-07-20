package fiber

import (
	"os"

	"github.com/wjojf/go-uber-fx/internal/api/http/server"
	"go.uber.org/fx"
)

var Options = fx.Provide(
	// Middleware
	fx.Annotate(
		server.WithMiddleware,
		fx.ParamTags(`group:"middlewares"`),
		fx.ResultTags(`group:"options"`),
	),

	// Static Files
	fx.Annotate(
		withStatic,
		fx.ResultTags(`group:"options"`),
	),
)

func withStatic() server.Option {
	cwd, _ := os.Getwd()
	return server.WithStatic("/static*", cwd+"/static")
}
