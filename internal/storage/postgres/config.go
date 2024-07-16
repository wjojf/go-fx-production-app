package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

func setupConfig(cfg config.Config) (*pgxpool.Config, error) {

	dbCfg, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	dbCfg.MaxConns = int32(cfg.DatabaseMaxConn)
	dbCfg.MaxConnIdleTime = time.Duration(cfg.DatabaseMaxIdle) * time.Second
	dbCfg.ConnConfig.ConnectTimeout = time.Duration(cfg.DatabaseTimeout) * time.Second

	return dbCfg, nil
}
