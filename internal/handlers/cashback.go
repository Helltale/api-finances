package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

// get one by id
func GetCashbackById(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCashbackById called", "method", r.Method)
	loggerFile.Info("GetCashbackById called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idCashbackStr := strings.TrimPrefix(r.URL.Path, "/cashback/id/")
	if idCashbackStr == "" {
		http.Error(w, u.JsonErrorResponse("id_cashback is required"), http.StatusBadRequest)
		return
	}

	idCashback, err := strconv.ParseInt(idCashbackStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_cashback"), http.StatusBadRequest)
		return
	}

	var foundCashback *models.Cashback
	if config.Mode == "debug" {
		for _, cashback := range debuging.Cashbacks {
			if cashback.GetIdCashback() == idCashback {
				foundCashback = cashback
				break
			}
		}
	}

	if foundCashback == nil {
		http.Error(w, u.JsonErrorResponse("Cashback not found"), http.StatusNotFound)
		return
	}

	cashbackJSON, err := foundCashback.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting cashback to JSON", "error", err)
		loggerFile.Error("Error converting cashback to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbackJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved cashback", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved cashback", "status", http.StatusOK)
}

// get by account id
func GetCashbacksByAccountId(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCashbacksByAccountId called", "method", r.Method)
	loggerFile.Info("GetCashbacksByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/cashback/account/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.Mode == "debug" {
		for _, cashback := range debuging.Cashbacks {
			if cashback.GetIdAccaunt() == idAccaunt {
				foundCashbacks = append(foundCashbacks, cashback)
			}
		}
	}

	if len(foundCashbacks) == 0 {
		http.Error(w, u.JsonErrorResponse("No cashbacks found for the account"), http.StatusNotFound)
		return
	}

	var cashbacksJSON []models.CashbackJSON
	for _, cashback := range foundCashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting cashback to JSON", "error", err)
			loggerFile.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved cashbacks for account", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved cashbacks for account", "status", http.StatusOK)
}

// get by bank name !!!
func GetCashbacksByBankName(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCashbacksByBankName called", "method", r.Method)
	loggerFile.Info("GetCashbacksByBankName called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	bankNameEncoded := strings.TrimPrefix(r.URL.Path, "/cashback/bank/")
	bankName, err := url.QueryUnescape(bankNameEncoded)
	if err != nil || bankName == "" {
		http.Error(w, u.JsonErrorResponse("bank_name is required"), http.StatusBadRequest)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.Mode == "debug" {
		for _, cashback := range debuging.Cashbacks {
			if strings.EqualFold(cashback.GetBankName(), bankName) {
				foundCashbacks = append(foundCashbacks, cashback)
			}
		}
	}

	if len(foundCashbacks) == 0 {
		http.Error(w, u.JsonErrorResponse("No cashbacks found for the bank"), http.StatusNotFound)
		return
	}

	var cashbacksJSON []models.CashbackJSON
	for _, cashback := range foundCashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting cashback to JSON", "error", err)
			loggerFile.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved cashbacks for bank", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved cashbacks for bank", "status", http.StatusOK)
}

// get by category !!!
func GetCashbacksByCategory(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCashbacksByCategory called", "method", r.Method)
	loggerFile.Info("GetCashbacksByCategory called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	category := strings.TrimPrefix(r.URL.Path, "/cashback/category/")
	if category == "" {
		http.Error(w, u.JsonErrorResponse("category is required"), http.StatusBadRequest)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.Mode == "debug" {
		for _, cashback := range debuging.Cashbacks {
			if strings.EqualFold(cashback.GetCategory(), category) {
				foundCashbacks = append(foundCashbacks, cashback)
			}
		}
	}

	if len(foundCashbacks) == 0 {
		http.Error(w, u.JsonErrorResponse("No cashbacks found for the category"), http.StatusNotFound)
		return
	}

	var cashbacksJSON []models.CashbackJSON
	for _, cashback := range foundCashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting cashback to JSON", "error", err)
			loggerFile.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved cashbacks for category", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved cashbacks for category", "status", http.StatusOK)
}

// get current
func GetCashbacksCurrent(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCurrentCashbacks called", "method", r.Method)
	loggerFile.Info("GetCurrentCashbacks called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.Mode == "debug" {
		for _, cashback := range debuging.Cashbacks {
			if cashback.GetDateActualTo().Equal(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)) {
				foundCashbacks = append(foundCashbacks, cashback)
			}
		}
	}

	if len(foundCashbacks) == 0 {
		http.Error(w, u.JsonErrorResponse("No current cashbacks found"), http.StatusNotFound)
		return
	}

	var cashbacksJSON []models.CashbackJSON
	for _, cashback := range foundCashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting cashback to JSON", "error", err)
			loggerFile.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved current cashbacks", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved current cashbacks", "status", http.StatusOK)
}

// todo остальные crud
