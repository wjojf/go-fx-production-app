package profiling

import (
	"context"
	"github.com/wjojf/go-uber-fx/internal/pkg/profiling"
	"go.uber.org/fx"
	"log/slog"
	"runtime/pprof"
)

func ProfilingHooks(lc fx.Lifecycle, log *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := profiling.Start("cpu.prof")
				if err != nil {
					log.Error("Error starting profiling", "error", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			pprof.StopCPUProfile()
			return nil
		},
	})
}
