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

func GetAllIncomesExpected(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllIncomesExpected called", "method", r.Method)
	loggerFile.Info("GetAllIncomesExpected called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesExpected []models.IncomeExpected
	if config.Mode == "debug" {
		incomesExpected = debuging.IncomesExpected
	} else {
		incomesExpected = []models.IncomeExpected{}
	}

	var incomesExpectedJSON []models.IncomeExpectedJSON
	for _, incomeExpected := range incomesExpected {
		incomeExpectedJSON, err := incomeExpected.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expected income to JSON", "error", err)
			loggerFile.Error("Error converting expected income to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expected income to JSON"), http.StatusInternalServerError)
			return
		}
		incomesExpectedJSON = append(incomesExpectedJSON, *incomeExpectedJSON)
	}

	if err := json.NewEncoder(w).Encode(incomesExpectedJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
}
