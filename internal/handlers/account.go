package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/helltale/api-finances/internal/models"
)

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var accountsJSON []models.AccountJSON
	for _, account := range accounts {
		accountJSON, err := account.ToJSON()
		if err != nil {
			http.Error(w, "Error converting income to JSON", http.StatusInternalServerError)
			return
		}
		accountsJSON = append(accountsJSON, *accountJSON)
	}

	if err := json.NewEncoder(w).Encode(accountsJSON); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
