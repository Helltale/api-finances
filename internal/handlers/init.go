package handlers

import (
	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/logger"
)

func Init(logger *logger.CombinedLogger, config *config.Config) {
	switch config.AppMode {
	case "debug":
		debugging.Init()
		logger.Info("run in debug mode")
	case "release":
		logger.Info("run in release mode")
	default:
		logger.Error("bad config.mode")
	}
}
