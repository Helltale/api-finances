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
	remains []models.Remain
)

func GetAllRemains(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllRemains called", "method", r.Method)
	loggerFile.Info("GetAllRemains called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var remainsJSON []models.RemainJSON
	if config.Mode == "debug" {
		remains = debuging.Remains
	} else {
		remains = []models.Remain{}
	}

	for _, remain := range remains {
		remainJSON, err := remain.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting remain to JSON", "error", err)
			loggerFile.Error("Error converting remain to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
			return
		}
		remainsJSON = append(remainsJSON, *remainJSON)
	}

	if err := json.NewEncoder(w).Encode(remainsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved remains", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved remains", "status", http.StatusOK)
}
