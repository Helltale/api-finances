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
func GetAllExpences(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllExpences called", "method", r.Method)
	loggerFile.Info("GetAllExpences called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var expences []*models.Expence
	if config.Mode == "debug" {
		expences = debuging.Expences
	} else {
		expences = []*models.Expence{}
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range expences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences", "status", http.StatusOK)
}

// todo получение записи по id

// todo получение записи по id группы траты

// todo получение записи по названию траты

// todo получение записи в промежутке времени fe: 	...between/2022-12-31/9999-12-31			- промежуток
//													...between?from=2022-12-31&to=9999-12-31	- или так, пока не решил

// todo получение всех постоянных и нет трат fe:	...every/1
//													.../ever/0

// todo получение записей по цене fe: 				...amount/between/1000/10000000
//													...amount/between?from=1000&to=10000000
//													...amount/more/10000000000
//													...amount/less/10000000000

//todo остальные crud
