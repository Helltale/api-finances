package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

var (
	incomes         []models.Income
	accounts        []models.Account
	incomesExpected []models.IncomeExpected
)

func GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesJSON []models.IncomeJSON
	for _, income := range incomes {
		incomeJSON, err := income.ToJSON()
		if err != nil {
			http.Error(w, "Error converting income to JSON", http.StatusInternalServerError)
			return
		}
		incomesJSON = append(incomesJSON, *incomeJSON)
	}

	if err := json.NewEncoder(w).Encode(incomesJSON); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
