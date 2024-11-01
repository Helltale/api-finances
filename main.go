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
	fileLogger, err := logger.NewFLogger(conf.AppFilelog)
	if err != nil {
		slogger.Error("ошибка создания logger", "config", conf, "error", err)
	}
	defer fileLogger.Close()

	logger := logger.NewCombinedLogger(slogger, fileLogger)

	handlers.Init(logger, conf)

	logger.Info("Server starting", "port", conf.AppPort)

	routers.Init(logger, conf)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", conf.AppPort), nil); err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}
