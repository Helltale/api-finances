package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func account(logger *logger.CombinedLogger, config config.Config) {
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetAll(w, r, logger, config)
	})
	http.HandleFunc("/account/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/account/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPost(w, r, logger, config)
	})
	http.HandleFunc("/account/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPut(w, r, logger, config)
	})
	http.HandleFunc("/account/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountDelete(w, r, logger, config)
	})
}
