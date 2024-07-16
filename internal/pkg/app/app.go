package app

import (
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func New() (*fx.App, error) {

	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	return fx.New(di.GetAppBundle(*cfg)), nil
}
