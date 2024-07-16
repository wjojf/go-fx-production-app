package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

var (
	pool *pgxpool.Pool
)

func NewPool(cfg config.Config) (*pgxpool.Pool, error) {

	if pool != nil {
		return pool, nil
	}

	dbCfg, err := setupConfig(cfg)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), dbCfg)
}
