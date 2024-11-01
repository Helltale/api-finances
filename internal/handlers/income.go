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
	u "github.com/helltale/api-finances/internal/utils"
)

//todo сделать группу методов с актуальными структурами

// get all
func IncomeGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {

	logger.Info("GetAllIncomes called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomes []*models.Income
	if config.AppMode == "debug" {
		incomes = debugging.Incomes
	} else {
		incomes = []*models.Income{}
	}

	response := make([]models.IncomeJSON, 0, len(incomes))
	for _, income := range incomes {
		incomeJSON, err := income.ToJSON()
		if err != nil {
			logger.Error("Error converting income to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting income to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *incomeJSON)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Info("GetAllIncomes called", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}
	logger.Info("Successfully retrieved incomes", "status", http.StatusOK)
}

// get one by id
func IncomeGetByIdIncome(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetIncomeById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idIncomeStr := strings.TrimPrefix(r.URL.Path, "/income/id/")
	if idIncomeStr == "" {
		http.Error(w, u.JsonErrorResponse("id_income is required"), http.StatusBadRequest)
		return
	}

	idIncome, err := strconv.ParseInt(idIncomeStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_income"), http.StatusBadRequest)
		return
	}

	var foundIncome *models.Income
	if config.AppMode == "debug" {
		for _, income := range debugging.Incomes {
			if income.GetIdIncome() == idIncome {
				foundIncome = income
				break
			}
		}
	}

	if foundIncome == nil {
		http.Error(w, u.JsonErrorResponse("Income not found"), http.StatusNotFound)
		return
	}

	incomeJSON, err := foundIncome.ToJSON()
	if err != nil {
		logger.Error("Error converting income to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error converting income to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incomeJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved income", "status", http.StatusOK)
}

// get all by person id
func IncomeGetByIdAccount(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetIncomesByAccountId called", "method", r.Method)
	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/income/account/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var incomes []*models.Income
	if config.AppMode == "debug" {
		incomes = debugging.Incomes
	} else {
		incomes = []*models.Income{}
	}

	var incomesByAccountId []models.IncomeJSON
	for _, income := range incomes {
		if income.GetIdAccaunt() == idAccaunt {
			incomeJSON, err := income.ToJSON()
			if err != nil {
				logger.Error("Error converting income to JSON", "error", err)
				http.Error(w, u.JsonErrorResponse("Error converting income to JSON"), http.StatusInternalServerError)
				return
			}
			incomesByAccountId = append(incomesByAccountId, *incomeJSON)
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
func IncomePost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PostIncome called", "method", r.Method)
	if r.Method != http.MethodPost {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newIncomeJSON models.IncomeJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newIncomeJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newIncome := &models.Income{}
	newIncome.SetIdIncome(newIncomeJSON.IdIncome)
	newIncome.SetIdAccaunt(newIncomeJSON.IdAccaunt)
	newIncome.SetIdIncomeExpected(newIncomeJSON.IdIncomeExpected)
	newIncome.SetAmount(newIncomeJSON.Amount)
	newIncome.SetExpectedAmount(newIncomeJSON.ExpectedAmount)
	newIncome.SetTypeIncome(newIncomeJSON.TypeIncome)
	newIncome.SetIncomeMonthMonth(newIncomeJSON.IncomeMonthMonth)
	newIncome.SetIncomeMonthDate(newIncomeJSON.IncomeMonthDate)
	newIncome.SetUpdBy(newIncomeJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newIncomeJSON.DateActualFrom); err == nil {
		newIncome.SetDateActualFrom(dateActualFrom)
	} else {
		logger.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newIncomeJSON.DateActualTo); err == nil {
		newIncome.SetDateActualTo(dateActualTo)
	} else {
		logger.Error("Error parsing DateActualTo", "error", err)
	}

	debugging.Incomes = append(debugging.Incomes, newIncome)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Income created successfully",
		"income":  newIncomeJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created income", "status", http.StatusCreated)
}

// update
func IncomePut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PutIncome called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Error("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/income/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idIncome, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedIncomeJSON models.IncomeJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedIncomeJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldIncome *models.Income
	for i, income := range debugging.Incomes {
		if income.GetIdIncome() == idIncome {
			oldIncome = income

			debugging.Incomes = append(debugging.Incomes[:i], debugging.Incomes[i+1:]...)
			break
		}
	}

	if oldIncome == nil {
		http.Error(w, u.JsonErrorResponse("Income not found"), http.StatusNotFound)
		return
	}

	oldIncomeJSON, err := oldIncome.ToJSON()
	if err != nil {
		logger.Error("Error converting old income to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old income"), http.StatusInternalServerError)
		return
	}

	newIncome := &models.Income{}
	newIncome.SetIdAccaunt(updatedIncomeJSON.IdAccaunt)
	newIncome.SetIdIncomeExpected(updatedIncomeJSON.IdIncomeExpected)
	newIncome.SetAmount(updatedIncomeJSON.Amount)
	newIncome.SetExpectedAmount(updatedIncomeJSON.ExpectedAmount)
	newIncome.SetTypeIncome(updatedIncomeJSON.TypeIncome)
	newIncome.SetIncomeMonthMonth(updatedIncomeJSON.IncomeMonthMonth)
	newIncome.SetIncomeMonthDate(updatedIncomeJSON.IncomeMonthDate)
	newIncome.SetUpdBy(updatedIncomeJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeJSON.DateActualFrom); err == nil {
		newIncome.SetDateActualFrom(dateActualFrom)
	} else {
		logger.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeJSON.DateActualTo); err == nil {
		newIncome.SetDateActualTo(dateActualTo)
	} else {
		logger.Error("Error parsing DateActualTo", "error", err)
	}

	debugging.Incomes = append(debugging.Incomes, newIncome)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":      "Income updated successfully",
		"index_income": idIncome,
		"old_income":   oldIncomeJSON,
		"new_income":   updatedIncomeJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated income", "status", http.StatusOK)
}

// delete
func IncomeDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("DeleteIncome called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/income/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idIncome, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldIncome *models.Income
	for i, income := range debugging.Incomes {
		if income.GetIdIncome() == idIncome {
			oldIncome = income

			debugging.Incomes = append(debugging.Incomes[:i], debugging.Incomes[i+1:]...)
			break
		}
	}

	if oldIncome == nil {
		http.Error(w, u.JsonErrorResponse("Income not found"), http.StatusNotFound)
		return
	}

	oldIncomeJSON, err := oldIncome.ToJSON()
	if err != nil {
		logger.Error("Error converting old income to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old income"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":      "Income deleted successfully",
		"index_income": idIncome,
		"old_income":   oldIncomeJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted income", "status", http.StatusOK)
}
