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

// get by bank name
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

// get by category
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

// create
func PostCashback(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PostCashback called", "method", r.Method)
	loggerFile.Info("PostCashback called", "method", r.Method)

	if r.Method != http.MethodPost {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newCashbackJSON models.CashbackJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCashbackJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newCashback := &models.Cashback{}
	newCashback.SetIdCashback(newCashbackJSON.IdCashback)
	newCashback.SetIdAccaunt(newCashbackJSON.IdAccaunt)
	newCashback.SetBankName(newCashbackJSON.BankName)
	newCashback.SetCategory(newCashbackJSON.Category)
	newCashback.SetPercent(newCashbackJSON.Percent)
	newCashback.SetUpdBy(newCashbackJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newCashbackJSON.DateActualFrom); err == nil {
		newCashback.SetDateActualFrom(dateActualFrom)
	}
	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newCashbackJSON.DateActualTo); err == nil {
		newCashback.SetDateActualTo(dateActualTo)
	}

	debuging.Cashbacks = append(debuging.Cashbacks, newCashback)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message":  "Cashback created successfully",
		"cashback": newCashbackJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully created cashback", "status", http.StatusCreated)
	loggerFile.Info("Successfully created cashback", "status", http.StatusCreated)
}

// update
func PutCashback(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PutCashback called", "method", r.Method)
	loggerFile.Info("PutCashback called", "method", r.Method)

	if r.Method != http.MethodPut {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/cashback/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idCashback, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedCashbackJSON models.CashbackJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedCashbackJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldCashback *models.Cashback
	for i, cashback := range debuging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			oldCashback = cashback

			debuging.Cashbacks = append(debuging.Cashbacks[:i], debuging.Cashbacks[i+1:]...)
			break
		}
	}

	if oldCashback == nil {
		http.Error(w, u.JsonErrorResponse("Cashback not found"), http.StatusNotFound)
		return
	}

	oldCashbackJSON, err := oldCashback.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old cashback to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old cashback"), http.StatusInternalServerError)
		return
	}

	newCashback := &models.Cashback{}
	newCashback.SetIdCashback(updatedCashbackJSON.IdCashback)
	newCashback.SetIdAccaunt(updatedCashbackJSON.IdAccaunt)
	newCashback.SetBankName(updatedCashbackJSON.BankName)
	newCashback.SetCategory(updatedCashbackJSON.Category)
	newCashback.SetPercent(updatedCashbackJSON.Percent)
	newCashback.SetUpdBy(updatedCashbackJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedCashbackJSON.DateActualFrom); err == nil {
		newCashback.SetDateActualFrom(dateActualFrom)
	} else {
		loggerConsole.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedCashbackJSON.DateActualTo); err == nil {
		newCashback.SetDateActualTo(dateActualTo)
	} else {
		loggerConsole.Error("Error parsing DateActualTo", "error", err)
	}

	debuging.Cashbacks = append(debuging.Cashbacks, newCashback)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":        "Cashback updated successfully",
		"index_cashback": idCashback,
		"old_cashback":   oldCashbackJSON,
		"new_cashback":   updatedCashbackJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully updated cashback", "status", http.StatusOK)
	loggerFile.Info("Successfully updated cashback", "status", http.StatusOK)
}

// delete
func DeleteCashback(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("DeleteCashback called", "method", r.Method)
	loggerFile.Info("DeleteCashback called", "method", r.Method)

	if r.Method != http.MethodDelete {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/cashback/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idCashback, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldCashback *models.Cashback
	for i, cashback := range debuging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			oldCashback = cashback

			// Удаление записи из среза
			debuging.Cashbacks = append(debuging.Cashbacks[:i], debuging.Cashbacks[i+1:]...)
			break
		}
	}

	if oldCashback == nil {
		http.Error(w, u.JsonErrorResponse("Cashback not found"), http.StatusNotFound)
		return
	}

	// Преобразуем старую структуру в JSON для ответа
	oldCashbackJSON, err := oldCashback.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old cashback to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old cashback"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":        "Cashback deleted successfully",
		"index_cashback": idCashback,
		"old_cashback":   oldCashbackJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully deleted cashback", "status", http.StatusOK)
	loggerFile.Info("Successfully deleted cashback", "status", http.StatusOK)
}
