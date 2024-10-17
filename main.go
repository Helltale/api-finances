package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/routers"
)

func main() {
	config.Init("config/config.yaml")

	handlers.Init(config.AppConf)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Server starting", "port", config.AppConf.APIPort)

	routers.Init(logger)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.AppConf.APIPort), nil); err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}
