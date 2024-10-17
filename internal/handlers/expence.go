package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

var (
	expences []models.Expence
)

func GetAllExpences(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger) {
	loggerConsole.Info("GetAllExpences called", "method", r.Method)
	loggerFile.Info("GetAllExpences called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, jsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var expencesJSON []models.ExpenceJSON
	for _, expence := range expences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, jsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, jsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences", "status", http.StatusOK)
}
