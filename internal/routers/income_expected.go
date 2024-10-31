package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func income_expected(logger *logger.CombinedLogger, config config.Config) {
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetAll(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedGetByIncomeExpectedId(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetByAccountId(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPost(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPut(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedDelete(w, r, logger, config)
	})
}
