package routers

import (
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/internal/handlers"
)

func Init(logger *slog.Logger) {
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomes(w, r, logger)
	})

	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomesExpected(w, r, logger)
	})

	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllAccounts(w, r, logger)
	})
}
