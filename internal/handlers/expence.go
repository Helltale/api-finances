package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/helltale/api-finances/config"
	debugging "github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/logger"
	"github.com/helltale/api-finances/internal/models"
	"github.com/helltale/api-finances/internal/services"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func ExpenceGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAllExpences called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	expenceService := services.NewExpenceService()
	expences := expenceService.GetAllExpences()

	if len(expences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range expences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved all expences", "status", http.StatusOK)
}

// get one by id
func ExpenceGetByIdExpence(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpenceById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, u.JsonErrorResponse("Expence ID is required"), http.StatusBadRequest)
		return
	}

	idStr := urlParts[4]
	idExpence, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid expence ID format"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	expence, err := expenceService.GetExpenceById(idExpence)
	if err != nil {
		logger.Error("Error fetching expence", "error", err)
		http.Error(w, u.JsonErrorResponse("Error fetching expence"), http.StatusInternalServerError)
		return
	}

	expenceJSON, err := expence.ToJSON()
	if err != nil {
		logger.Error("Error converting expence to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expenceJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expence by ID", "status", http.StatusOK)

}

// get all by group id
func ExpenceGetByIdGroup(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByGroup called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, u.JsonErrorResponse("Group is required"), http.StatusBadRequest)
		return
	}

	group := urlParts[4]

	expenceService := services.NewExpenceService()
	expences, err := expenceService.GetExpencesByGroup(group)
	if err != nil {
		logger.Error("Error fetching expences", "error", err)
		http.Error(w, u.JsonErrorResponse("Error fetching expences"), http.StatusInternalServerError)
		return
	}

	if len(expences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found for the specified group"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range expences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by group", "status", http.StatusOK)
}

// get all by title
func ExpenceGetByTitle(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByTitle called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	titleExpence := strings.TrimPrefix(r.URL.Path, "/expence/title/")
	if titleExpence == "" {
		http.Error(w, u.JsonErrorResponse("title_expence is required"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByTitle(titleExpence)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by title", "status", http.StatusOK)
}

// get all by date range
func ExpenceGetByDateBetween(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByDateRange called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 4 {
		http.Error(w, u.JsonErrorResponse("Both start and end dates are required"), http.StatusBadRequest)
		return
	}

	startDateStr := urlParts[4]
	endDateStr := urlParts[5]

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid start date format"), http.StatusBadRequest)
		return
	}

	var endDate time.Time
	if endDateStr == "9999-12-31" {
		endDate = time.Now()
	} else {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, u.JsonErrorResponse("Invalid end date format"), http.StatusBadRequest)
			return
		}
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by date range", "status", http.StatusOK)
}

// get all by repeat type
func ExpenceGetByRepeat(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByRepeatType called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 3 {
		http.Error(w, u.JsonErrorResponse("Repeat type is required"), http.StatusBadRequest)
		return
	}

	repeatStr := urlParts[3]
	repeat, err := strconv.ParseInt(repeatStr, 10, 8)
	if err != nil || (repeat != 0 && repeat != 1) {
		http.Error(w, u.JsonErrorResponse("Invalid repeat type, must be 0 or 1"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByRepeat(int8(repeat))
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by repeat type", "status", http.StatusOK)
}

// get all by amount range
func ExpenceGetByAmountBetween(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByAmountRange called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 6 {
		http.Error(w, u.JsonErrorResponse("Both min and max amounts are required"), http.StatusBadRequest)
		return
	}

	minAmountStr := urlParts[4]
	maxAmountStr := urlParts[5]

	minAmount, err := strconv.ParseFloat(minAmountStr, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid min amount format"), http.StatusBadRequest)
		return
	}

	maxAmount, err := strconv.ParseFloat(maxAmountStr, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid max amount format"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByAmountRange(minAmount, maxAmount)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by amount range", "status", http.StatusOK)
}

// get all by max amount
func ExpenceGetByAmountLess(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByMaxAmount called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, u.JsonErrorResponse("Max amount is required"), http.StatusBadRequest)
		return
	}

	maxAmountStr := urlParts[4]

	maxAmount, err := strconv.ParseFloat(maxAmountStr, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid max amount format"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByMaxAmount(maxAmount)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by max amount", "status", http.StatusOK)
}

// get all by min amount
func ExpencesGetByAmountMore(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetExpencesByMinAmount called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, u.JsonErrorResponse("Min amount is required"), http.StatusBadRequest)
		return
	}

	minAmountStr := urlParts[4]

	minAmount, err := strconv.ParseFloat(minAmountStr, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid min amount format"), http.StatusBadRequest)
		return
	}

	expenceService := services.NewExpenceService()
	foundExpences, err := expenceService.GetExpencesByMinAmount(minAmount)
	if err != nil {
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			logger.Error("Error converting expence to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved expences by min amount", "status", http.StatusOK)
}

// create
func ExpencePost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PostExpence called", "method", r.Method)

	if r.Method != http.MethodPost {
		logger.Info("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newExpenceJSON models.ExpenceJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newExpenceJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newExpence := &models.Expence{}
	newExpence.SetIdExpence(newExpenceJSON.IdExpence)
	newExpence.SetGroupExpence(newExpenceJSON.GroupExpence)
	newExpence.SetTitleExpence(newExpenceJSON.TitleExpence)
	newExpence.SetDescriptionExpence(newExpenceJSON.DescriptionExpence)
	newExpence.SetRepeat(newExpenceJSON.Repeat)
	newExpence.SetAmount(newExpenceJSON.Amount)
	newExpence.SetUpdBy(newExpenceJSON.UpdBy)

	if date, err := time.Parse("2006-01-02T15:04:05Z", newExpenceJSON.Date); err == nil {
		newExpence.SetDate(date)
	}
	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newExpenceJSON.DateActualFrom); err == nil {
		newExpence.SetDateActualFrom(dateActualFrom)
	}
	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newExpenceJSON.DateActualTo); err == nil {
		newExpence.SetDateActualTo(dateActualTo)
	}

	debugging.Expences = append(debugging.Expences, newExpence)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Expence created successfully",
		"expence": newExpenceJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created expence", "status", http.StatusCreated)
}

// update
func ExpencePut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PutExpence called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/expence/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idExpence, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedExpenceJSON models.ExpenceJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedExpenceJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldExpence *models.Expence
	for i, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			oldExpence = expence

			debugging.Expences = append(debugging.Expences[:i], debugging.Expences[i+1:]...)
			break
		}
	}

	if oldExpence == nil {
		http.Error(w, u.JsonErrorResponse("Expence not found"), http.StatusNotFound)
		return
	}

	oldExpenceJSON, err := oldExpence.ToJSON()
	if err != nil {
		logger.Error("Error converting old expence to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old expence"), http.StatusInternalServerError)
		return
	}

	newExpence := &models.Expence{}
	newExpence.SetIdExpence(updatedExpenceJSON.IdExpence)
	newExpence.SetGroupExpence(updatedExpenceJSON.GroupExpence)
	newExpence.SetTitleExpence(updatedExpenceJSON.TitleExpence)
	newExpence.SetDescriptionExpence(updatedExpenceJSON.DescriptionExpence)
	newExpence.SetRepeat(updatedExpenceJSON.Repeat)
	newExpence.SetAmount(updatedExpenceJSON.Amount)
	newExpence.SetUpdBy(updatedExpenceJSON.UpdBy)

	if date, err := time.Parse("2006-01-02T15:04:05Z", updatedExpenceJSON.Date); err == nil {
		newExpence.SetDate(date)
	} else {
		logger.Error("Error parsing Date", "error", err)
	}

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedExpenceJSON.DateActualFrom); err == nil {
		newExpence.SetDateActualFrom(dateActualFrom)
	} else {
		logger.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedExpenceJSON.DateActualTo); err == nil {
		newExpence.SetDateActualTo(dateActualTo)
	} else {
		logger.Error("Error parsing DateActualTo", "error", err)
	}

	debugging.Expences = append(debugging.Expences, newExpence)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":       "Expence updated successfully",
		"index_expence": idExpence,
		"old_expence":   oldExpenceJSON,
		"new_expence":   updatedExpenceJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated expence", "status", http.StatusOK)
}

// delete
func ExpenceDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("DeleteExpence called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/expence/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idExpence, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldExpence *models.Expence
	for i, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			oldExpence = expence

			debugging.Expences = append(debugging.Expences[:i], debugging.Expences[i+1:]...)
			break
		}
	}

	if oldExpence == nil {
		http.Error(w, u.JsonErrorResponse("Expence not found"), http.StatusNotFound)
		return
	}

	oldExpenceJSON, err := oldExpence.ToJSON()
	if err != nil {
		logger.Error("Error converting old expence to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old expence"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":       "Expence deleted successfully",
		"index_expence": idExpence,
		"old_expence":   oldExpenceJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted expence", "status", http.StatusOK)
}
