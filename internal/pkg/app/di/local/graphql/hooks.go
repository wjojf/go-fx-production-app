package graphql

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func Hooks(lc fx.Lifecycle, cfg config.Config, log *slog.Logger, mux *http.ServeMux) {

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("Starting GraphQL Server")
			go http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), mux)
			return nil
		},
	})

}
