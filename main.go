package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
	"github.com/helltale/api-finances/internal/routers"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v\n", err)
	}

	slogger := logger.NewSLogger()
	fileLogger, err := logger.NewFLogger(conf.FilepathLog)
	if err != nil {
		slogger.Error("Ошибка создания FileLogger", "error", err)
	}
	defer fileLogger.Close()

	logger := logger.NewCombinedLogger(slogger, fileLogger)

	handlers.Init(logger, config.AppConf)

	logger.Info("Server starting", "port", config.AppConf.APIPort)

	routers.Init(logger, config.AppConf)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.AppConf.APIPort), nil); err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}
