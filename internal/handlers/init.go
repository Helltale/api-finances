package handlers

import (
	"log/slog"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debuging"
)

func Init(config config.Config) {
	switch config.Mode {
	case "debug":
		debuging.Init()
		slog.Info("run in debug mode")
	case "release":
		slog.Info("run in release mode")
	default:
		slog.Error("bad config.mode")
	}
}