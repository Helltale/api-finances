package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debuging"
	"github.com/helltale/api-finances/internal/models"
	u "github.com/helltale/api-finances/internal/utils"
)

var (
	incomes         []models.Income
	accounts        []models.Account
	incomesExpected []models.IncomeExpected
)

func GetAllIncomes(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllIncomes called", "method", r.Method)
	loggerFile.Info("GetAllIncomes called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesJSON []models.IncomeJSON
	if config.Mode == "debug" {
		incomes = debuging.Incomes
	} else {
		incomes = []models.Income{}
	}

	for _, income := range incomes {
		incomeJSON, err := income.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting income to JSON", "error", err)
			loggerFile.Error("Error converting income to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting income to JSON"), http.StatusInternalServerError)
			return
		}
		incomesJSON = append(incomesJSON, *incomeJSON)
	}

	if err := json.NewEncoder(w).Encode(incomesJSON); err != nil {
		loggerConsole.Info("GetAllIncomes called", "method", r.Method)
		loggerFile.Info("GetAllIncomes called", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved incomes", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved incomes", "status", http.StatusOK)

}
