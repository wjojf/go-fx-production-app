package di

import (
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/dev"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/local"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func GetAppBundle(cfg config.Config) fx.Option {
	switch cfg.Environment {
	case config.EnvLocal:
		return local.Bundle(cfg)
	case config.EnvDev:
		return dev.Bundle(cfg)
	case config.EnvProd:
		return prod.Bundle(cfg)
	default:
		return prod.Bundle(cfg)
	}
}
