package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func cashback(logger *logger.CombinedLogger, config *config.Config) {
	http.HandleFunc("/cashback/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetAll(w, r, logger, config)
	})
	http.HandleFunc("/cashback/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdCashback(w, r, logger, config)
	})
	http.HandleFunc("/cashback/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/cashback/bank/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByBankName(w, r, logger, config)
	})
	http.HandleFunc("/cashback/category/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByCategory(w, r, logger, config)
	})
	http.HandleFunc("/cashback/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetCurrent(w, r, logger, config)
	})
	http.HandleFunc("/cashback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPost(w, r, logger, config)
	})
	http.HandleFunc("/cashback/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPut(w, r, logger, config)
	})
	http.HandleFunc("/cashback/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackDelete(w, r, logger, config)
	})
}
