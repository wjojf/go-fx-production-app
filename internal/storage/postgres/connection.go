package postgres

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

var (
	pool *pgxpool.Pool
)

func NewPool(cfg config.Config, log *slog.Logger) (*pgxpool.Pool, error) {

	if pool != nil {
		return pool, nil
	}

	dbCfg, err := setupConfig(cfg, log)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), dbCfg)
}
