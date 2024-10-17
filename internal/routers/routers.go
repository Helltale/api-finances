package routers

import (
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/internal/handlers"
)

func Init(loggerConsole *slog.Logger, loggerFile *slog.Logger) {
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomes(w, r, loggerConsole, loggerFile)
	})
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomesExpected(w, r, loggerConsole, loggerFile)
	})
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllAccounts(w, r, loggerConsole, loggerFile)
	})
	http.HandleFunc("/expence/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllExpences(w, r, loggerConsole, loggerFile)
	})
}
