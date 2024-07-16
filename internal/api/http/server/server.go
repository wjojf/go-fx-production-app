package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

func New(cfg config.Config, fiberCfg fiber.Config) *fiber.App {
	return fiber.New(fiberCfg)
}
