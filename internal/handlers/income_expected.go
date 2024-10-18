package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debuging"
	"github.com/helltale/api-finances/internal/models"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func GetIncomesExpectedAll(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllIncomesExpected called", "method", r.Method)
	loggerFile.Info("GetAllIncomesExpected called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var incomesExpected []*models.IncomeExpected
	if config.Mode == "debug" {
		incomesExpected = debuging.IncomesExpected
	} else {
		incomesExpected = []*models.IncomeExpected{}
	}

	response := make([]models.IncomeExpectedJSON, 0, len(incomesExpected))
	for _, incomeExpected := range incomesExpected {
		jsonIncomeExpected, err := incomeExpected.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting to JSON", "error", err)
			loggerFile.Error("Error converting to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *jsonIncomeExpected)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expected incomes", "status", http.StatusOK)
}

// get one by id
func GetIncomesExpectedById(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetIncomeExpectedById called", "method", r.Method)
	loggerFile.Info("GetIncomeExpectedById called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
	if config.Mode == "debug" {
		for _, incomeExpected := range debuging.IncomesExpected {
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
		loggerConsole.Error("Error converting expected income to JSON", "error", err)
		loggerFile.Error("Error converting expected income to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting expected income to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(incomeExpectedJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved expected income", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved expected income", "status", http.StatusOK)
}

// get all by person id
func GetIncomesExpectedByAccountId(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetIncomesByAccountId called", "method", r.Method)
	loggerFile.Info("GetIncomesByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
	if config.Mode == "debug" {
		incomesExpected = debuging.IncomesExpected
	} else {
		incomesExpected = []*models.IncomeExpected{}
	}

	var incomesByAccountId []models.IncomeExpectedJSON
	for _, incomeExpected := range incomesExpected {
		if incomeExpected.GetIdAccaunt() == idAccaunt {
			incomeExpectedJSON, err := incomeExpected.ToJSON()
			if err != nil {
				loggerConsole.Error("Error converting expected income to JSON", "error", err)
				loggerFile.Error("Error converting expected income to JSON", "error", err)

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
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved incomes for account ID", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved incomes for account ID", "status", http.StatusOK)
}

// create
func PostIncomeExpected(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PostIncomeExpected called", "method", r.Method)
	loggerFile.Info("PostIncomeExpected called", "method", r.Method)

	if r.Method != http.MethodPost {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newIncomeExpectedJSON models.IncomeExpectedJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newIncomeExpectedJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

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

	debuging.IncomesExpected = append(debuging.IncomesExpected, newIncomeExpected)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message":         "Income expected created successfully",
		"income_expected": newIncomeExpectedJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully created income expected", "status", http.StatusCreated)
	loggerFile.Info("Successfully created income expected", "status", http.StatusCreated)
}

// update
func PutIncomeExpected(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PutIncomeExpected called", "method", r.Method)
	loggerFile.Info("PutIncomeExpected called", "method", r.Method)

	if r.Method != http.MethodPut {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		loggerConsole.Warn("ID not provided")
		loggerFile.Warn("ID not provided")

		http.Error(w, u.JsonErrorResponse("ID not provided"), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		loggerConsole.Error("Error parsing ID", "error", err)
		loggerFile.Error("Error parsing ID", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid ID format"), http.StatusBadRequest)
		return
	}

	var updatedIncomeExpectedJSON models.IncomeExpectedJSON
	decoder := json.NewDecoder(r.Body)
	fmt.Println("-----------------------------------", r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&updatedIncomeExpectedJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	loggerConsole.Info("Decoded JSON", "data", updatedIncomeExpectedJSON)
	loggerFile.Info("Decoded JSON", "data", updatedIncomeExpectedJSON)

	existingIncomeExpected, err := GetIncomeExpectedById(id)
	if err != nil {
		loggerConsole.Error("Error fetching income expected", "error", err)
		loggerFile.Error("Error fetching income expected", "error", err)

		http.Error(w, u.JsonErrorResponse("Income expected not found"), http.StatusNotFound)
		return
	}

	existingIncomeExpected.SetIdAccaunt(updatedIncomeExpectedJSON.IdAccaunt)
	existingIncomeExpected.SetIdIncomeEx(updatedIncomeExpectedJSON.IdIncomeEx)
	existingIncomeExpected.SetAmount(updatedIncomeExpectedJSON.Amount)
	existingIncomeExpected.SetTypeIncome(updatedIncomeExpectedJSON.TypeIncome)
	existingIncomeExpected.SetIncomeMonthDate(updatedIncomeExpectedJSON.IncomeMonthDate)
	existingIncomeExpected.SetUpdBy(updatedIncomeExpectedJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeExpectedJSON.DateActualFrom); err == nil {
		existingIncomeExpected.SetDateActualFrom(dateActualFrom)
	} else {
		loggerConsole.Warn("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedIncomeExpectedJSON.DateActualTo); err == nil {
		existingIncomeExpected.SetDateActualTo(dateActualTo)
	} else {
		loggerConsole.Warn("Error parsing DateActualTo", "error", err)
	}

	for i, incomeExpected := range debuging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == id {
			debuging.IncomesExpected[i] = existingIncomeExpected
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":         "Income expected updated successfully",
		"income_expected": existingIncomeExpected,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully updated income expected", "status", http.StatusOK)
	loggerFile.Info("Successfully updated income expected", "status", http.StatusOK)
}

func GetIncomeExpectedById(idIncomeEx int64) (*models.IncomeExpected, error) {
	for _, incomeExpected := range debuging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == idIncomeEx {
			return incomeExpected, nil
		}
	}
	return nil, errors.New("income expected not found")
}

func SaveIncomeExpected(incomeExpected *models.IncomeExpected) error {
	debuging.IncomesExpected = append(debuging.IncomesExpected, incomeExpected)
	return nil
}
