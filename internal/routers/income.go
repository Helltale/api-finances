package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func income(logger *logger.CombinedLogger, config config.Config) {
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetAll(w, r, logger, config)
	})
	http.HandleFunc("/income/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdIncome(w, r, logger, config)
	})
	http.HandleFunc("/income/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/income/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePost(w, r, logger, config)
	})
	http.HandleFunc("/income/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePut(w, r, logger, config)
	})
	http.HandleFunc("/income/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeDelete(w, r, logger, config)
	})
}
