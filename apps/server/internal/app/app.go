package app

import (
	"flakeforge/fin/config"
	"flakeforge/fin/pkg/logger"
)

func Run(cfg *config.Config) {
  l := logger.New(cfg.Logger.Level)
}
