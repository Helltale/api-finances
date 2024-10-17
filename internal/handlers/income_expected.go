package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

func GetAllIncomesExpected(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	logger.Info("GetAllIncomesExpected called", "method", r.Method)
	if r.Method != http.MethodGet {
		logger.Warn("Method not allowed", "method", r.Method)
		http.Error(w, jsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesExpectedJSON []models.IncomeExpectedJSON
	for _, incomeExpected := range incomesExpected {
		incomeExpectedJSON, err := incomeExpected.ToJSON()
		if err != nil {
			logger.Error("Error converting expected income to JSON", "error", err)
			http.Error(w, jsonErrorResponse("Error converting expected income to JSON"), http.StatusInternalServerError)
			return
		}
		incomesExpectedJSON = append(incomesExpectedJSON, *incomeExpectedJSON)
	}

	if err := json.NewEncoder(w).Encode(incomesExpectedJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, jsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
}
