package routers

import (
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
)

func Init(loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomes(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomesExpected(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllAccounts(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllExpences(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/remain/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllRemains(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllGoals(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllCashbacks(w, r, loggerConsole, loggerFile, config)
	})
}
