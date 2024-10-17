package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/helltale/api-finances/models"
)

var incomes []models.Income

func GetAllIncomes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesJSON []models.IncomeJSON
	for _, income := range incomes {
		incomesJSON = append(incomesJSON, income.ToJSON())
	}

	json.NewEncoder(w).Encode(incomesJSON)
}
