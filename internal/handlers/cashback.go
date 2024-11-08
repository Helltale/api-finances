package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/logger"
	"github.com/helltale/api-finances/internal/models"
	"github.com/helltale/api-finances/internal/services"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func CashbackGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAllCashbacks called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var cashbacks []*models.Cashback
	if config.AppMode == "debug" {
		cashbacks = debugging.Cashbacks
	} else {
		cashbacks = []*models.Cashback{}
	}

	response := make([]models.CashbackJSON, 0, len(cashbacks))
	for _, cashback := range cashbacks {
		cashbackJSON, err := cashback.ToJSON()
		if err != nil {
			logger.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *cashbackJSON)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved cashbacks", "status", http.StatusOK)
}

// get one by id
func CashbackGetByIdCashback(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetCashbackById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

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
	if config.AppMode == "debug" {
		for _, cashback := range debugging.Cashbacks {
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
		logger.Error("Error converting cashback to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbackJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved cashback", "status", http.StatusOK)
}

// get by account id
func CashbackGetByIdAccount(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetCashbacksByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

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
	if config.AppMode == "debug" {
		for _, cashback := range debugging.Cashbacks {
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
			logger.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved cashbacks for account", "status", http.StatusOK)
}

// get by bank name
func CashbackGetByBankName(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetCashbacksByBankName called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

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
	if config.AppMode == "debug" {
		for _, cashback := range debugging.Cashbacks {
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
			logger.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved cashbacks for bank", "status", http.StatusOK)
}

// get by category
func CashbackGetByCategory(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetCashbacksByCategory called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	category := strings.TrimPrefix(r.URL.Path, "/cashback/category/")
	if category == "" {
		http.Error(w, u.JsonErrorResponse("category is required"), http.StatusBadRequest)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.AppMode == "debug" {
		for _, cashback := range debugging.Cashbacks {
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
			logger.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved cashbacks for category", "status", http.StatusOK)
}

// get current
func CashbackGetCurrent(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetCurrentCashbacks called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var foundCashbacks []*models.Cashback
	if config.AppMode == "debug" {
		for _, cashback := range debugging.Cashbacks {
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
			logger.Error("Error converting cashback to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting cashback to JSON"), http.StatusInternalServerError)
			return
		}
		cashbacksJSON = append(cashbacksJSON, *cashbackJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cashbacksJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved current cashbacks", "status", http.StatusOK)
}

// create
func CashbackPost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("CashbackAdd called", "method", r.Method)

	if r.Method != http.MethodPost {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newCashbackJSON models.CashbackJSON
	if err := json.NewDecoder(r.Body).Decode(&newCashbackJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	// Преобразование JSON структуры в модель Cashback
	newCashback := &models.Cashback{}
	newCashback.SetIdCashback(newCashbackJSON.IdCashback)
	newCashback.SetIdAccaunt(newCashbackJSON.IdAccaunt)
	newCashback.SetBankName(newCashbackJSON.BankName)
	newCashback.SetCategory(newCashbackJSON.Category)
	newCashback.SetPercent(newCashbackJSON.Percent)
	newCashback.SetUpdBy(newCashbackJSON.UpdBy)
	dateFrom, _ := time.Parse("2006-01-02", newCashbackJSON.DateActualFrom)
	dateTo, _ := time.Parse("2006-01-02", newCashbackJSON.DateActualTo)
	newCashback.SetDateActualFrom(dateFrom)
	newCashback.SetDateActualTo(dateTo)

	cashbackService := services.NewCashbackService()
	if err := cashbackService.AddNewCashback(newCashback); err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message":     "Cashback added successfully",
		"id_cashback": newCashbackJSON.IdCashback,
	}
	json.NewEncoder(w).Encode(response)
}

// update
func CashbackPut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("CashbackUpdate called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var updatedCashbackJSON models.CashbackJSON
	if err := json.NewDecoder(r.Body).Decode(&updatedCashbackJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	updatedCashback := &models.Cashback{}
	updatedCashback.SetIdCashback(updatedCashbackJSON.IdCashback)
	updatedCashback.SetIdAccaunt(updatedCashbackJSON.IdAccaunt)
	updatedCashback.SetBankName(updatedCashbackJSON.BankName)
	updatedCashback.SetCategory(updatedCashbackJSON.Category)
	updatedCashback.SetPercent(updatedCashbackJSON.Percent)
	updatedCashback.SetUpdBy(updatedCashbackJSON.UpdBy)
	dateFrom, _ := time.Parse("2006-01-02", updatedCashbackJSON.DateActualFrom)
	dateTo, _ := time.Parse("2006-01-02", updatedCashbackJSON.DateActualTo)
	updatedCashback.SetDateActualFrom(dateFrom)
	updatedCashback.SetDateActualTo(dateTo)

	cashbackService := services.NewCashbackService()
	if _, err := cashbackService.UpdateCashback(updatedCashback); err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":     "Cashback updated successfully",
		"id_cashback": updatedCashbackJSON.IdCashback,
	}
	json.NewEncoder(w).Encode(response)
}

// update + tohistory
func CashbackUpdateWithHistory(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger) {
	logger.Info("CashbackUpdateWithHistory called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var updatedCashbackJSON models.CashbackJSON
	if err := json.NewDecoder(r.Body).Decode(&updatedCashbackJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	// Создаем новую актуальную запись
	newCashback := &models.Cashback{}
	newCashback.SetIdCashback(updatedCashbackJSON.IdCashback)
	newCashback.SetIdAccaunt(updatedCashbackJSON.IdAccaunt)
	newCashback.SetBankName(updatedCashbackJSON.BankName)
	newCashback.SetCategory(updatedCashbackJSON.Category)
	newCashback.SetPercent(updatedCashbackJSON.Percent)
	newCashback.SetUpdBy(updatedCashbackJSON.UpdBy)

	cashbackService := services.NewCashbackService()
	oldCashback, err := cashbackService.UpdateHistoryCashback(updatedCashbackJSON.IdCashback, newCashback)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":      "Cashback updated with history successfully",
		"id_cashback":  updatedCashbackJSON.IdCashback,
		"old_cashback": oldCashback,
	}
	json.NewEncoder(w).Encode(response)
}

// delete
func CashbackDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("CashbackDelete called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	idStr := strings.TrimPrefix(urlPath, "/cashback/delete/")
	idCashback, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid cashback ID format"), http.StatusBadRequest)
		return
	}

	cashbackService := services.NewCashbackService()
	deletedCashback, err := cashbackService.DeleteCashback(idCashback)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":      "Cashback deleted successfully",
		"id_cashback":  idCashback,
		"deleted_data": deletedCashback,
	}
	json.NewEncoder(w).Encode(response)
}

// delete + restore
func CashbackDeleteAndRestore(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger) {
	logger.Info("CashbackDeleteAndRestore called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	idStr := strings.TrimPrefix(urlPath, "/cashback/delete-and-restore/")
	idCashback, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid cashback ID format"), http.StatusBadRequest)
		return
	}

	cashbackService := services.NewCashbackService()
	restoredCashback, err := cashbackService.DeleteAndRestorePreviousCashback(idCashback)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":       "Cashback deleted and restored successfully",
		"id_cashback":   idCashback,
		"restored_data": restoredCashback,
	}
	json.NewEncoder(w).Encode(response)
}
