package logging

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres/logging/adapters"
)

func AddLogging(cfg config.Config, log *slog.Logger, dbCfg *pgxpool.Config) {

	l := adapters.NewSlogTracer(log)

	switch cfg.Environment {

	case config.EnvProd:
		dbCfg.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   l,
			LogLevel: tracelog.LogLevelError,
		}

	case config.EnvDev, config.EnvLocal:
		dbCfg.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   l,
			LogLevel: tracelog.LogLevelDebug,
		}

	}
}
