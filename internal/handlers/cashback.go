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

// get all
func GetAllCashbacks(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllCashbacks called", "method", r.Method)
	loggerFile.Info("GetAllCashbacks called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var cashbacks []*models.Cashback
	if config.Mode == "debug" {
		cashbacks = debuging.Cashbacks
	} else {
		cashbacks = []*models.Cashback{}
	}

	response := make([]models.CashbackJSON, 0, len(cashbacks))
	for _, cashback := range cashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting cashback to JSON", "error", err)
			loggerFile.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *cashbackJSON)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved cashbacks", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved cashbacks", "status", http.StatusOK)
}

// todo получение всех кэшбеков по id аккаунта

// todo получение всех кэшбеков по названию банка

// todo получение всех кэшбеков по категории

// todo получение всех актуальных кэшбеков

// todo остальные crud
