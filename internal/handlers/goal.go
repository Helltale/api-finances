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
	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/models"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func GoalGetAll(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllGoals called", "method", r.Method)
	loggerFile.Info("GetAllGoals called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var goals []*models.Goal
	if config.Mode == "debug" {
		goals = debugging.Goals
	} else {
		goals = []*models.Goal{}
	}

	response := make([]models.GoalJSON, 0, len(goals))
	for _, goal := range goals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *goalJSON)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals", "status", http.StatusOK)
}

// get one by id
func GoalGetByIdGoal(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalById called", "method", r.Method)
	loggerFile.Info("GetGoalById called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idGoalStr := strings.TrimPrefix(r.URL.Path, "/goal/id/")
	if idGoalStr == "" {
		http.Error(w, u.JsonErrorResponse("id_goal is required"), http.StatusBadRequest)
		return
	}

	idGoal, err := strconv.ParseInt(idGoalStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_goal"), http.StatusBadRequest)
		return
	}

	var foundGoal *models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetIdGoal() == idGoal {
				foundGoal = goal
				break
			}
		}
	}

	if foundGoal == nil {
		http.Error(w, u.JsonErrorResponse("Goal not found"), http.StatusNotFound)
		return
	}

	goalJSON, err := foundGoal.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting goal to JSON", "error", err)
		loggerFile.Error("Error converting goal to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goal", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goal", "status", http.StatusOK)
}

// get all by account id
func GoalGetByIdAccount(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalsByAccountId called", "method", r.Method)
	loggerFile.Info("GetGoalsByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/goal/account/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var goals []*models.Goal
	if config.Mode == "debug" {
		goals = debugging.Goals
	} else {
		goals = []*models.Goal{}
	}

	var goalsByAccountId []models.GoalJSON
	for _, goal := range goals {
		if goal.GetIdAccaunt() == idAccaunt {
			goalJSON, err := goal.ToJSON()
			if err != nil {
				loggerConsole.Error("Error converting goal to JSON", "error", err)
				loggerFile.Error("Error converting goal to JSON", "error", err)

				http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
				return
			}
			goalsByAccountId = append(goalsByAccountId, *goalJSON)
		}
	}

	if len(goalsByAccountId) == 0 {
		http.Error(w, u.JsonErrorResponse("No goals found for the given account ID"), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsByAccountId); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals for account ID", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals for account ID", "status", http.StatusOK)
}

// get all by date range
func GoalGetByDateBetween(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalsByDateRange called", "method", r.Method)
	loggerFile.Info("GetGoalsByDateRange called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 5 {
		http.Error(w, u.JsonErrorResponse("Both start and end dates are required"), http.StatusBadRequest)
		return
	}

	startDateStr := urlParts[4]
	endDateStr := urlParts[5]
	fmt.Println(startDateStr)
	fmt.Println(endDateStr)

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

	var foundGoals []*models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetDate().After(startDate) && goal.GetDate().Before(endDate) {
				foundGoals = append(foundGoals, goal)
			}
		}
	}

	if len(foundGoals) == 0 {
		http.Error(w, u.JsonErrorResponse("No goals found in the specified date range"), http.StatusNotFound)
		return
	}

	var goalsJSON []models.GoalJSON
	for _, goal := range foundGoals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals by date range", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals by date range", "status", http.StatusOK)
}

// get current
func GoalGetCurrent(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetCurrentGoals called", "method", r.Method)
	loggerFile.Info("GetCurrentGoals called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var foundGoals []*models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetDateActualTo().Equal(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)) {
				foundGoals = append(foundGoals, goal)
			}
		}
	}

	if len(foundGoals) == 0 {
		http.Error(w, u.JsonErrorResponse("No current goals found"), http.StatusNotFound)
		return
	}

	var goalsJSON []models.GoalJSON
	for _, goal := range foundGoals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved current goals", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved current goals", "status", http.StatusOK)
}

// get all by amount range
func GoalGetByAmountBetween(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalsByAmountRange called", "method", r.Method)
	loggerFile.Info("GetGoalsByAmountRange called", "method", r.Method)

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

	var foundGoals []*models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetAmount() >= minAmount && goal.GetAmount() <= maxAmount {
				foundGoals = append(foundGoals, goal)
			}
		}
	}

	if len(foundGoals) == 0 {
		http.Error(w, u.JsonErrorResponse("No goals found in the specified amount range"), http.StatusNotFound)
		return
	}

	var goalsJSON []models.GoalJSON
	for _, goal := range foundGoals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals by amount range", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals by amount range", "status", http.StatusOK)
}

// get all by max amount
func GoalGetByAmountLess(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalsByMaxAmount called", "method", r.Method)
	loggerFile.Info("GetGoalsByMaxAmount called", "method", r.Method)

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

	var foundGoals []*models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetAmount() < maxAmount {
				foundGoals = append(foundGoals, goal)
			}
		}
	}

	if len(foundGoals) == 0 {
		http.Error(w, u.JsonErrorResponse("No goals found below the specified amount"), http.StatusNotFound)
		return
	}

	var goalsJSON []models.GoalJSON
	for _, goal := range foundGoals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals by max amount", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals by max amount", "status", http.StatusOK)
}

// get all by min amount
func GoalGetByAmountMore(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetGoalsByMinAmount called", "method", r.Method)
	loggerFile.Info("GetGoalsByMinAmount called", "method", r.Method)

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

	var foundGoals []*models.Goal
	if config.Mode == "debug" {
		for _, goal := range debugging.Goals {
			if goal.GetAmount() > minAmount {
				foundGoals = append(foundGoals, goal)
			}
		}
	}

	if len(foundGoals) == 0 {
		http.Error(w, u.JsonErrorResponse("No goals found above the specified amount"), http.StatusNotFound)
		return
	}

	var goalsJSON []models.GoalJSON
	for _, goal := range foundGoals {
		goalJSON, err := goal.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting goal to JSON", "error", err)
			loggerFile.Error("Error converting goal to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting goal to JSON"), http.StatusInternalServerError)
			return
		}
		goalsJSON = append(goalsJSON, *goalJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goalsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved goals by min amount", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved goals by min amount", "status", http.StatusOK)
}

// create
func GoalPost(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PostGoal called", "method", r.Method)
	loggerFile.Info("PostGoal called", "method", r.Method)

	if r.Method != http.MethodPost {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newGoalJSON models.GoalJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newGoalJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newGoal := &models.Goal{}
	newGoal.SetIdGoal(newGoalJSON.IdGoal)
	newGoal.SetIdAccaunt(newGoalJSON.IdAccaunt)
	newGoal.SetAmount(newGoalJSON.Amount)
	newGoal.SetUpdBy(newGoalJSON.UpdBy)

	if date, err := time.Parse("2006-01-02T15:04:05Z", newGoalJSON.Date); err == nil {
		newGoal.SetDate(date)
	}
	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newGoalJSON.DateActualFrom); err == nil {
		newGoal.SetDateActualFrom(dateActualFrom)
	}
	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newGoalJSON.DateActualTo); err == nil {
		newGoal.SetDateActualTo(dateActualTo)
	}

	debugging.Goals = append(debugging.Goals, newGoal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Goal created successfully",
		"goal":    newGoalJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully created goal", "status", http.StatusCreated)
	loggerFile.Info("Successfully created goal", "status", http.StatusCreated)
}

// update
func GoalPut(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PutGoal called", "method", r.Method)
	loggerFile.Info("PutGoal called", "method", r.Method)

	if r.Method != http.MethodPut {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/goal/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idGoal, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedGoalJSON models.GoalJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedGoalJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldGoal *models.Goal
	for i, goal := range debugging.Goals {
		if goal.GetIdGoal() == idGoal {
			oldGoal = goal

			debugging.Goals = append(debugging.Goals[:i], debugging.Goals[i+1:]...)
			break
		}
	}

	if oldGoal == nil {
		http.Error(w, u.JsonErrorResponse("Goal not found"), http.StatusNotFound)
		return
	}

	oldGoalJSON, err := oldGoal.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old goal to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old goal"), http.StatusInternalServerError)
		return
	}

	newGoal := &models.Goal{}
	newGoal.SetIdGoal(updatedGoalJSON.IdGoal)
	newGoal.SetIdAccaunt(updatedGoalJSON.IdAccaunt)
	newGoal.SetAmount(updatedGoalJSON.Amount)
	newGoal.SetUpdBy(updatedGoalJSON.UpdBy)

	if date, err := time.Parse("2006-01-02T15:04:05Z", updatedGoalJSON.Date); err == nil {
		newGoal.SetDate(date)
	} else {
		loggerConsole.Error("Error parsing Date", "error", err)
	}

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedGoalJSON.DateActualFrom); err == nil {
		newGoal.SetDateActualFrom(dateActualFrom)
	} else {
		loggerConsole.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedGoalJSON.DateActualTo); err == nil {
		newGoal.SetDateActualTo(dateActualTo)
	} else {
		loggerConsole.Error("Error parsing DateActualTo", "error", err)
	}

	debugging.Goals = append(debugging.Goals, newGoal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":    "Goal updated successfully",
		"index_goal": idGoal,
		"old_goal":   oldGoalJSON,
		"new_goal":   updatedGoalJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully updated goal", "status", http.StatusOK)
	loggerFile.Info("Successfully updated goal", "status", http.StatusOK)
}

// delete
func GoalDelete(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("DeleteGoal called", "method", r.Method)
	loggerFile.Info("DeleteGoal called", "method", r.Method)

	if r.Method != http.MethodDelete {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/goal/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idGoal, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldGoal *models.Goal
	for i, goal := range debugging.Goals {
		if goal.GetIdGoal() == idGoal {
			oldGoal = goal

			debugging.Goals = append(debugging.Goals[:i], debugging.Goals[i+1:]...)
			break
		}
	}

	if oldGoal == nil {
		http.Error(w, u.JsonErrorResponse("Goal not found"), http.StatusNotFound)
		return
	}

	oldGoalJSON, err := oldGoal.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old goal to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old goal"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":    "Goal deleted successfully",
		"index_goal": idGoal,
		"old_goal":   oldGoalJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully deleted goal", "status", http.StatusOK)
	loggerFile.Info("Successfully deleted goal", "status", http.StatusOK)
}
