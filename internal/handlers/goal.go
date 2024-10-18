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
	goals []models.Goal
)

func GetAllGoals(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllGoals called", "method", r.Method)
	loggerFile.Info("GetAllGoals called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var goalsJSON []models.GoalJSON
	if config.Mode == "debug" {
		goals = debuging.Goals
	} else {
		goals = []models.Goal{}
	}

	for _, goal := range goals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals", "status", http.StatusOK)
}
