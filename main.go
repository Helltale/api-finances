package main

import (
	"fmt"
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/routers"
)

func main() {
	config.Init()

	handlers.Init()
	routers.Init()

	fmt.Printf("info: server start on localhost:%s\n", config.AppConf.APIPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.AppConf.APIPort), nil)
}
