package di

import (
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/local"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func GetAppBundle(cfg config.Config) fx.Option {
	switch cfg.Environment {
	case config.EnvLocal:
		return local.Bundle(cfg)
	default:
		panic("Unknown environment")
	}
}
