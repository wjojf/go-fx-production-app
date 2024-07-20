package local

import (
	"fmt"

	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func Bundle(cfg config.Config) fx.Option {
	return fx.Options(
		fx.Invoke(
			func() {
				fmt.Println("Hello, World!")
			},
		),
	)
}
