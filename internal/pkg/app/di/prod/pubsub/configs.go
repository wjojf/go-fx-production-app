package pubsub

import (
	"context"
	"os"

	"github.com/wjojf/go-uber-fx/internal/events/pubsub"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

func GetClientOpts(cfg config.Config) pubsub.ClientOpts {

	cwd, _ := os.Getwd()

	return pubsub.ClientOpts{
		Ctx:         context.Background(),
		ProjectID:   cfg.GoogleProjectID,
		KeyFilepath: cwd + "/keys/pubsub-prod.json",
	}
}

func GetSubscriptionConfigurator(cfg config.Config) pubsub.SubscriptionConfigurator {
	return pubsub.DefaultSubscriptionConfigurator
}

func GetContextFunc() pubsub.ContextFunc {
	return pubsub.DefaultContextFunc
}

func GetOperationFunc() pubsub.OperationFunc {
	return pubsub.DefaultOperationGenerator
}
