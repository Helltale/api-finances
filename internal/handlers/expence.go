package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/helltale/api-finances/config"
	debugging "github.com/helltale/api-finances/internal/debuging"
	"github.com/helltale/api-finances/internal/models"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func ExpenceGetAll(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllExpences called", "method", r.Method)
	loggerFile.Info("GetAllExpences called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var expences []*models.Expence
	if config.Mode == "debug" {
		expences = debugging.Expences
	} else {
		expences = []*models.Expence{}
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range expences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences", "status", http.StatusOK)
}

// get one by id
func ExpenceGetByIdExpence(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpenceById called", "method", r.Method)
	loggerFile.Info("GetExpenceById called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idExpenceStr := strings.TrimPrefix(r.URL.Path, "/expence/id/")
	if idExpenceStr == "" {
		http.Error(w, u.JsonErrorResponse("id_expence is required"), http.StatusBadRequest)
		return
	}

	idExpence, err := strconv.ParseInt(idExpenceStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_expence"), http.StatusBadRequest)
		return
	}

	var foundExpence *models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetIdExpence() == idExpence {
				foundExpence = expence
				break
			}
		}
	}

	if foundExpence == nil {
		http.Error(w, u.JsonErrorResponse("Expence not found"), http.StatusNotFound)
		return
	}

	expenceJSON, err := foundExpence.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting expence to JSON", "error", err)
		loggerFile.Error("Error converting expence to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expenceJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expence", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expence", "status", http.StatusOK)
}

// get all by group id
func ExpenceGetByIdGroup(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByGroupId called", "method", r.Method)
	loggerFile.Info("GetExpencesByGroupId called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	groupExpence := strings.TrimPrefix(r.URL.Path, "/expence/group/")
	if groupExpence == "" {
		http.Error(w, u.JsonErrorResponse("group_expence is required"), http.StatusBadRequest)
		return
	}

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetGroupExpence() == groupExpence {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found for the specified group"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by group", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by group", "status", http.StatusOK)
}

// get all by title
func ExpenceGetByTitle(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByTitle called", "method", r.Method)
	loggerFile.Info("GetExpencesByTitle called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	titleExpence := strings.TrimPrefix(r.URL.Path, "/expence/title/")
	if titleExpence == "" {
		http.Error(w, u.JsonErrorResponse("title_expence is required"), http.StatusBadRequest)
		return
	}

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetTitleExpence() == titleExpence {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found for the specified title"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by title", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by title", "status", http.StatusOK)
}

// get all by date range
func ExpenceGetByDateBetween(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByDateRange called", "method", r.Method)
	loggerFile.Info("GetExpencesByDateRange called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetDate().After(startDate) && expence.GetDate().Before(endDate) {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found in the specified date range"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by date range", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by date range", "status", http.StatusOK)
}

// get all by repeat type
func ExpenceGetByRepeat(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByRepeatType called", "method", r.Method)
	loggerFile.Info("GetExpencesByRepeatType called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	fmt.Println(urlParts)
	if len(urlParts) < 2 {
		http.Error(w, u.JsonErrorResponse("Repeat type is required"), http.StatusBadRequest)
		return
	}

	repeatStr := urlParts[3]
	repeat, err := strconv.ParseInt(repeatStr, 10, 8)
	if err != nil || (repeat != 0 && repeat != 1) {
		http.Error(w, u.JsonErrorResponse("Invalid repeat type, must be 0 or 1"), http.StatusBadRequest)
		return
	}

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetRepeat() == int8(repeat) {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found for the specified repeat type"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by repeat type", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by repeat type", "status", http.StatusOK)
}

// get all by amount range
func ExpenceGetByAmountBetween(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByAmountRange called", "method", r.Method)
	loggerFile.Info("GetExpencesByAmountRange called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 6 {
		http.Error(w, u.JsonErrorResponse("Both min and max amounts are required"), http.StatusBadRequest)
		return
	}

	fmt.Println(urlParts)

	minAmountStr := urlParts[4]
	fmt.Println(minAmountStr)
	maxAmountStr := urlParts[5]
	fmt.Println(maxAmountStr)

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

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetAmount() >= minAmount && expence.GetAmount() <= maxAmount {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found in the specified amount range"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by amount range", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by amount range", "status", http.StatusOK)
}

// get all by max amount
func ExpenceGetByAmountLess(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByMaxAmount called", "method", r.Method)
	loggerFile.Info("GetExpencesByMaxAmount called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetAmount() < maxAmount {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found below the specified amount"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by max amount", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by max amount", "status", http.StatusOK)
}

// get all by min amount
func ExpencesGetByAmountMore(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetExpencesByMinAmount called", "method", r.Method)
	loggerFile.Info("GetExpencesByMinAmount called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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

	var foundExpences []*models.Expence
	if config.Mode == "debug" {
		for _, expence := range debugging.Expences {
			if expence.GetAmount() > minAmount {
				foundExpences = append(foundExpences, expence)
			}
		}
	}

	if len(foundExpences) == 0 {
		http.Error(w, u.JsonErrorResponse("No expences found above the specified amount"), http.StatusNotFound)
		return
	}

	var expencesJSON []models.ExpenceJSON
	for _, expence := range foundExpences {
		expenceJSON, err := expence.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting expence to JSON", "error", err)
			loggerFile.Error("Error converting expence to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting expence to JSON"), http.StatusInternalServerError)
			return
		}
		expencesJSON = append(expencesJSON, *expenceJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(expencesJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expences by min amount", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expences by min amount", "status", http.StatusOK)
}

// create
func ExpencePost(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PostExpence called", "method", r.Method)
	loggerFile.Info("PostExpence called", "method", r.Method)

	if r.Method != http.MethodPost {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newExpenceJSON models.ExpenceJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newExpenceJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

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
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully created expence", "status", http.StatusCreated)
	loggerFile.Info("Successfully created expence", "status", http.StatusCreated)
}

// update
func ExpencePut(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PutExpence called", "method", r.Method)
	loggerFile.Info("PutExpence called", "method", r.Method)

	if r.Method != http.MethodPut {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

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
		loggerConsole.Error("Error converting old expence to JSON", "error", err)
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
		loggerConsole.Error("Error parsing Date", "error", err)
	}

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedExpenceJSON.DateActualFrom); err == nil {
		newExpence.SetDateActualFrom(dateActualFrom)
	} else {
		loggerConsole.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedExpenceJSON.DateActualTo); err == nil {
		newExpence.SetDateActualTo(dateActualTo)
	} else {
		loggerConsole.Error("Error parsing DateActualTo", "error", err)
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
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully updated expence", "status", http.StatusOK)
	loggerFile.Info("Successfully updated expence", "status", http.StatusOK)
}

// delete
func ExpenceDelete(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("DeleteExpence called", "method", r.Method)
	loggerFile.Info("DeleteExpence called", "method", r.Method)

	if r.Method != http.MethodDelete {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
		loggerConsole.Error("Error converting old expence to JSON", "error", err)
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
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully deleted expence", "status", http.StatusOK)
	loggerFile.Info("Successfully deleted expence", "status", http.StatusOK)
}
