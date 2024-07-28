package handlers

import (
	"github.com/wjojf/go-uber-fx/internal/events/pubsub/handlers/user"
	"go.uber.org/fx"
)

// This module repsents a collection of
// all pub/sub handlers available for use within the application.
var Module = fx.Module(
	"pub/sub-handlers",

	fx.Provide(user.NewVerifyHandler),
)
