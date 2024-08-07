package postgres

import (
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/internal/storage/postgres/logging"
)

func setupConfig(cfg config.Config, log *slog.Logger) (*pgxpool.Config, error) {

	dbCfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	dbCfg.MaxConns = int32(cfg.DatabaseMaxConn)
	dbCfg.MaxConnIdleTime = time.Duration(cfg.DatabaseMaxIdle) * time.Second
	dbCfg.ConnConfig.ConnectTimeout = time.Duration(cfg.DatabaseTimeout) * time.Second

	logging.AddLogging(cfg, log, dbCfg)

	return dbCfg, nil
}
