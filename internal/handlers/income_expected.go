package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

func GetAllIncomesExpected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesExpectedJSON []models.IncomeExpectedJSON
	for _, incomeExpected := range incomesExpected {
		incomeExpectedJSON, err := incomeExpected.ToJSON()
		if err != nil {
			http.Error(w, "Error converting income to JSON", http.StatusInternalServerError)
			return
		}
		incomesExpectedJSON = append(incomesExpectedJSON, *incomeExpectedJSON)
	}

	if err := json.NewEncoder(w).Encode(incomesExpectedJSON); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
