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

	logFile, err := os.OpenFile(config.AppConf.FilepathLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	handlers.Init(config.AppConf)

	loggerConsole := slog.New(slog.NewTextHandler(os.Stdout, nil))
	loggerFile := slog.New(slog.NewTextHandler(logFile, nil))

	loggerConsole.Info("Server starting", "port", config.AppConf.APIPort)
	loggerFile.Info("Server starting", "port", config.AppConf.APIPort)

	routers.Init(loggerConsole, loggerFile)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.AppConf.APIPort), nil); err != nil {
		loggerConsole.Error("Server failed to start", "error", err)
		loggerFile.Error("Server failed to start", "error", err)
	}
}
