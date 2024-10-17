package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

func GetAllAccounts(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger) {
	loggerConsole.Info("GetAllAccounts called", "method", r.Method)
	loggerFile.Info("GetAllAccounts called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, jsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var accountsJSON []models.AccountJSON
	for _, account := range accounts {
		accountJSON, err := account.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting account to JSON", "error", err)
			loggerFile.Error("Error converting account to JSON", "error", err)

			http.Error(w, jsonErrorResponse("Error converting account to JSON"), http.StatusInternalServerError)
			return
		}
		accountsJSON = append(accountsJSON, *accountJSON)
	}

	if err := json.NewEncoder(w).Encode(accountsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, jsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved accounts", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved accounts", "status", http.StatusOK)

}
