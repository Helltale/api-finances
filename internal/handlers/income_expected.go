package handlers

import (
	"encoding/json"
	"net/http"
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

//todo сделать группу методов с актуальными структурами

// get all
func IncomesExpectedGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAllIncomesExpected called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesExpected []*models.IncomeExpected
	if config.AppMode == "debug" {
		incomesExpected = debugging.IncomesExpected
	} else {
		incomesExpected = []*models.IncomeExpected{}
	}

	response := make([]models.IncomeExpectedJSON, 0, len(incomesExpected))
	for _, incomeExpected := range incomesExpected {
		jsonIncomeExpected, err := incomeExpected.ToJSON()
		if err != nil {
			logger.Error("Error converting to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *jsonIncomeExpected)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
}

// get one by id
func IncomeExpectedGetByIncomeExpectedId(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetIncomeExpectedById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idIncomeExStr := strings.TrimPrefix(r.URL.Path, "/income_expected/id/")
	if idIncomeExStr == "" {
		http.Error(w, u.JsonErrorResponse("id_income_ex is required"), http.StatusBadRequest)
		return
	}

	idIncomeEx, err := strconv.ParseInt(idIncomeExStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_income_ex"), http.StatusBadRequest)
		return
	}

	var foundIncomeExpected *models.IncomeExpected
	if config.AppMode == "debug" {
		for _, incomeExpected := range debugging.IncomesExpected {
			if incomeExpected.GetIdIncomeEx() == idIncomeEx {
				foundIncomeExpected = incomeExpected
				break
			}
		}
	}

	if foundIncomeExpected == nil {
		http.Error(w, u.JsonErrorResponse("Income not found"), http.StatusNotFound)
		return
	}

	incomeExpectedJSON, err := foundIncomeExpected.ToJSON()
	if err != nil {
		logger.Error("Error converting expected income to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error converting expected income to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incomeExpectedJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expected income", "status", http.StatusOK)
}

// get all by person id
func IncomesExpectedGetByAccountId(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetIncomesByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/income_expected/account/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var incomesExpected []*models.IncomeExpected
	if config.AppMode == "debug" {
		incomesExpected = debugging.IncomesExpected
	} else {
		incomesExpected = []*models.IncomeExpected{}
	}

	var incomesByAccountId []models.IncomeExpectedJSON
	for _, incomeExpected := range incomesExpected {
		if incomeExpected.GetIdAccaunt() == idAccaunt {
			incomeExpectedJSON, err := incomeExpected.ToJSON()
			if err != nil {
				logger.Error("Error converting expected income to JSON", "error", err)
				http.Error(w, u.JsonErrorResponse("Error converting expected income to JSON"), http.StatusInternalServerError)
				return
			}
			incomesByAccountId = append(incomesByAccountId, *incomeExpectedJSON)
		}
	}

	if len(incomesByAccountId) == 0 {
		http.Error(w, u.JsonErrorResponse("No incomes found for the given account ID"), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incomesByAccountId); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved incomes for account ID", "status", http.StatusOK)
}

// create
func IncomeExpectedPost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("IncomeExpectedPost called", "method", r.Method)

	if r.Method != http.MethodPost {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newIncomeExpectedJSON models.IncomeExpectedJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newIncomeExpectedJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newIncomeExpected := &models.IncomeExpected{}
	newIncomeExpected.SetIdAccaunt(newIncomeExpectedJSON.IdAccaunt)
	newIncomeExpected.SetIdIncomeEx(newIncomeExpectedJSON.IdIncomeEx)
	newIncomeExpected.SetAmount(newIncomeExpectedJSON.Amount)
	newIncomeExpected.SetTypeIncome(newIncomeExpectedJSON.TypeIncome)
	newIncomeExpected.SetIncomeMonthDate(newIncomeExpectedJSON.IncomeMonthDate)
	newIncomeExpected.SetUpdBy(newIncomeExpectedJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newIncomeExpectedJSON.DateActualFrom); err == nil {
		newIncomeExpected.SetDateActualFrom(dateActualFrom)
	}
	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newIncomeExpectedJSON.DateActualTo); err == nil {
		newIncomeExpected.SetDateActualTo(dateActualTo)
	}

	// Use the service to add the new income expected
	incomeExpectedService := services.NewIncomeExpectedService()
	if err := incomeExpectedService.AddNewIncomeExpected(newIncomeExpected); err != nil {
		logger.Error("Error adding income expected", "error", err)
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusConflict)
		return
	}

	// Response with success message and created income expected data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message":         "Income expected created successfully",
		"income_expected": newIncomeExpectedJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created income expected", "status", http.StatusCreated)
}

// update
func IncomeExpectedPut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PutIncomeExpected called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body to get updated income expected data
	var updatedIncomeExpectedJSON models.IncomeExpectedJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedIncomeExpectedJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	// Convert JSON struct to model struct
	newIncomeExpected := &models.IncomeExpected{}
	newIncomeExpected.SetIdAccaunt(updatedIncomeExpectedJSON.IdAccaunt)
	newIncomeExpected.SetIdIncomeEx(updatedIncomeExpectedJSON.IdIncomeEx)
	newIncomeExpected.SetAmount(updatedIncomeExpectedJSON.Amount)
	newIncomeExpected.SetTypeIncome(updatedIncomeExpectedJSON.TypeIncome)
	newIncomeExpected.SetIncomeMonthDate(updatedIncomeExpectedJSON.IncomeMonthDate)
	newIncomeExpected.SetUpdBy(updatedIncomeExpectedJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeExpectedJSON.DateActualFrom); err == nil {
		newIncomeExpected.SetDateActualFrom(dateActualFrom)
	} else {
		logger.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeExpectedJSON.DateActualTo); err == nil {
		newIncomeExpected.SetDateActualTo(dateActualTo)
	} else {
		logger.Error("Error parsing DateActualTo", "error", err)
	}

	incomeExpectedService := services.NewIncomeExpectedService()
	oldIncomeExpected, err := incomeExpectedService.UpdateIncomeExpected(newIncomeExpected)
	if err != nil {
		logger.Error("Income expected not found", "error", err)
		http.Error(w, u.JsonErrorResponse("Income expected not found"), http.StatusNotFound)
		return
	}

	oldIncomeExpectedJSON, err := oldIncomeExpected.ToJSON()
	if err != nil {
		logger.Error("Error converting old income expected to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old income expected"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":             "Income expected updated successfully",
		"old_income_expected": oldIncomeExpectedJSON,
		"new_income_expected": updatedIncomeExpectedJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated income expected", "status", http.StatusOK)
}

// update + tohistory
func IncomeExpectedVersionUpdate(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("IncomeExpectedVersionUpdate called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем `idIncomeEx` из URL
	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/income_expected/version_update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idIncomeEx, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid ID format"), http.StatusBadRequest)
		return
	}

	// Декодируем JSON с новыми данными
	var updatedIncomeExpectedJSON models.IncomeExpectedJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedIncomeExpectedJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	// Создаем новую сущность `IncomeExpected` с новыми данными
	newIncomeExpected := &models.IncomeExpected{}
	newIncomeExpected.SetIdAccaunt(updatedIncomeExpectedJSON.IdAccaunt)
	newIncomeExpected.SetIdIncomeEx(updatedIncomeExpectedJSON.IdIncomeEx)
	newIncomeExpected.SetAmount(updatedIncomeExpectedJSON.Amount)
	newIncomeExpected.SetTypeIncome(updatedIncomeExpectedJSON.TypeIncome)
	newIncomeExpected.SetIncomeMonthDate(updatedIncomeExpectedJSON.IncomeMonthDate)
	newIncomeExpected.SetUpdBy(updatedIncomeExpectedJSON.UpdBy)

	incomeService := services.NewIncomeExpectedService()

	oldIncomeExpected, err := incomeService.UpdateHistoryIncomeExpected(idIncomeEx, newIncomeExpected)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	oldIncomeExpectedJSON, err := oldIncomeExpected.ToJSON()
	if err != nil {
		logger.Error("Error converting old income expected to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old income expected"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":             "Income expected updated with new version",
		"old_income_expected": oldIncomeExpectedJSON,
		"new_income_expected": updatedIncomeExpectedJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated income expected with new version", "status", http.StatusOK)
}

// delete
func IncomeExpectedDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("DeleteIncomeExpected called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var deleteIncomeExpectedJSON struct {
		IdIncomeEx int64 `json:"id_income_ex"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&deleteIncomeExpectedJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	incomeExpectedService := services.NewIncomeExpectedService()

	oldIncomeExpected, err := incomeExpectedService.DeleteIncomeExpected(deleteIncomeExpectedJSON.IdIncomeEx)
	if err != nil {
		logger.Error("Income expected not found", "error", err)
		http.Error(w, u.JsonErrorResponse("Income expected not found"), http.StatusNotFound)
		return
	}

	oldIncomeExpectedJSON, err := oldIncomeExpected.ToJSON()
	if err != nil {
		logger.Error("Error converting old income expected to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old income expected"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":             "Income expected deleted successfully",
		"id_income_expected":  deleteIncomeExpectedJSON.IdIncomeEx,
		"old_income_expected": oldIncomeExpectedJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted income expected", "status", http.StatusOK)
}

// delete + restore
func IncomeExpectedDeleteRestore(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("IncomeExpectedDeleteRestore called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/income_expected/delete_restore/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idIncomeEx, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid ID format"), http.StatusBadRequest)
		return
	}

	incomeService := services.NewIncomeExpectedService()

	// delete current and restore historical
	restoredRecord, err := incomeService.DeleteAndRestorePreviousIncomeExpexted(idIncomeEx)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	restoredRecordJSON, err := restoredRecord.ToJSON()
	if err != nil {
		logger.Error("Error converting restored income expected to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing restored income expected"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":                  "Current income expected deleted and last historical version restored",
		"restored_income_expected": restoredRecordJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted current record and restored last historical version", "status", http.StatusOK)
}
