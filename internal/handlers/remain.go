package handlers

import (
	"encoding/json"
	"fmt"
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

// get all
func RemainGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAllRemains called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var remains []*models.Remain
	if config.AppMode == "debug" {
		remains = debugging.Remains
	} else {
		remains = []*models.Remain{}
	}

	response := make([]models.RemainJSON, 0, len(remains))
	for _, remain := range remains {
		remainJSON, err := remain.ToJSON()
		if err != nil {
			logger.Error("Error converting remain to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
			return
		}
		response = append(response, *remainJSON)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved remains", "status", http.StatusOK)
}

// get one by id
func RemainGetByIdRemain(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetRemainById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idRemainsStr := strings.TrimPrefix(r.URL.Path, "/remain/id/")
	if idRemainsStr == "" {
		http.Error(w, u.JsonErrorResponse("id_remains is required"), http.StatusBadRequest)
		return
	}

	idRemains, err := strconv.ParseInt(idRemainsStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_remains"), http.StatusBadRequest)
		return
	}

	var foundRemain *models.Remain
	if config.AppMode == "debug" {
		for _, remain := range debugging.Remains {
			if remain.GetIdRemains() == idRemains {
				foundRemain = remain
				break
			}
		}
	}

	if foundRemain == nil {
		http.Error(w, u.JsonErrorResponse("Remain not found"), http.StatusNotFound)
		return
	}

	remainJSON, err := foundRemain.ToJSON()
	if err != nil {
		logger.Error("Error converting remain to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(remainJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved remain", "status", http.StatusOK)
}

// get all by account id
func RemainGetByIdAccount(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetRemainsByAccountId called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/remain/account/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var remains []*models.Remain
	if config.AppMode == "debug" {
		remains = debugging.Remains
	} else {
		remains = []*models.Remain{}
	}

	var remainsByAccountId []models.RemainJSON
	for _, remain := range remains {
		if remain.GetIdAccaunt() == idAccaunt {
			remainJSON, err := remain.ToJSON()
			if err != nil {
				logger.Error("Error converting remain to JSON", "error", err)

				http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
				return
			}
			remainsByAccountId = append(remainsByAccountId, *remainJSON)
		}
	}

	if len(remainsByAccountId) == 0 {
		http.Error(w, u.JsonErrorResponse("No remains found for the given account ID"), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(remainsByAccountId); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved remains for account ID", "status", http.StatusOK)
}

// get last entry by remain id
func RemainGetByIdLastEntry(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetLastRemainEntryById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	remainIdStr := strings.TrimPrefix(r.URL.Path, "/remain/last/id/")
	if remainIdStr == "" {
		http.Error(w, u.JsonErrorResponse("remain_id is required"), http.StatusBadRequest)
		return
	}

	remainId, err := strconv.ParseInt(remainIdStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid remain_id"), http.StatusBadRequest)
		return
	}

	var lastRemain *models.Remain
	if config.AppMode == "debug" {
		for _, remain := range debugging.Remains {
			if remain.GetIdRemains() == remainId {
				if remain.GetDateActualTo().Format("2006-01-02") == "9999-12-31" {
					lastRemain = remain
					break
				}
			}
		}
	}

	if lastRemain == nil {
		http.Error(w, u.JsonErrorResponse("Remain entry not found or not active"), http.StatusNotFound)
		return
	}

	remainJSON, err := lastRemain.ToJSON()
	if err != nil {
		logger.Error("Error converting remain to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(remainJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved last remain entry", "status", http.StatusOK)
}

// get entries by dateActualFrom between two dates
func RemainGetByDateBetween(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetRemainByDateBetween called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 6 {
		http.Error(w, u.JsonErrorResponse("Invalid URL format"), http.StatusBadRequest)
		return
	}

	startDateStr := urlParts[4]
	fmt.Println(startDateStr)
	endDateStr := urlParts[5]
	fmt.Println(endDateStr)

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid start date format"), http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid end date format"), http.StatusBadRequest)
		return
	}

	var foundRemains []*models.Remain
	if config.AppMode == "debug" {
		for _, remain := range debugging.Remains {
			dateActualFrom := remain.GetDateActualFrom()
			if dateActualFrom.After(startDate) && dateActualFrom.Before(endDate) {
				foundRemains = append(foundRemains, remain)
			}
		}
	}

	if len(foundRemains) == 0 {
		http.Error(w, u.JsonErrorResponse("No remains found in the specified date range"), http.StatusNotFound)
		return
	}

	var remainsJSON []models.RemainJSON
	for _, remain := range foundRemains {
		remainJSON, err := remain.ToJSON()
		if err != nil {
			logger.Error("Error converting remain to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting remain to JSON"), http.StatusInternalServerError)
			return
		}
		remainsJSON = append(remainsJSON, *remainJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(remainsJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved remains in date range", "status", http.StatusOK)
}

// create
func RemainPost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PostRemain called", "method", r.Method)

	if r.Method != http.MethodPost {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newRemainJSON models.RemainJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newRemainJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newRemain := &models.Remain{}
	newRemain.SetIdRemains(newRemainJSON.IdRemains)
	newRemain.SetIdAccaunt(newRemainJSON.IdAccaunt)
	newRemain.SetAmount(newRemainJSON.Amount)
	newRemain.SetLastUpdateAmount(newRemainJSON.LastUpdateAmount)
	newRemain.SetLastUpdateId(newRemainJSON.LastUpdateId)
	newRemain.SetLastUpdateGroup(newRemainJSON.LastUpdateGroup)
	newRemain.SetUpdBy(newRemainJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", newRemainJSON.DateActualFrom); err == nil {
		newRemain.SetDateActualFrom(dateActualFrom)
	}
	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", newRemainJSON.DateActualTo); err == nil {
		newRemain.SetDateActualTo(dateActualTo)
	}

	debugging.Remains = append(debugging.Remains, newRemain)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Remain created successfully",
		"remain":  newRemainJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created remain", "status", http.StatusCreated)
}

// update
func RemainPut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PutRemain called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/remain/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idRemain, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedRemainJSON models.RemainJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedRemainJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldRemain *models.Remain
	for i, remain := range debugging.Remains {
		if remain.GetIdRemains() == idRemain {
			oldRemain = remain

			debugging.Remains = append(debugging.Remains[:i], debugging.Remains[i+1:]...)
			break
		}
	}

	if oldRemain == nil {
		http.Error(w, u.JsonErrorResponse("Remain not found"), http.StatusNotFound)
		return
	}

	oldRemainJSON, err := oldRemain.ToJSON()
	if err != nil {
		logger.Error("Error converting old remain to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old remain"), http.StatusInternalServerError)
		return
	}

	newRemain := &models.Remain{}
	newRemain.SetIdRemains(updatedRemainJSON.IdRemains)
	newRemain.SetIdAccaunt(updatedRemainJSON.IdAccaunt)
	newRemain.SetAmount(updatedRemainJSON.Amount)
	newRemain.SetLastUpdateAmount(updatedRemainJSON.LastUpdateAmount)
	newRemain.SetLastUpdateId(updatedRemainJSON.LastUpdateId)
	newRemain.SetLastUpdateGroup(updatedRemainJSON.LastUpdateGroup)
	newRemain.SetUpdBy(updatedRemainJSON.UpdBy)

	if dateActualFrom, err := time.Parse("2006-01-02T15:04:05Z", updatedRemainJSON.DateActualFrom); err == nil {
		newRemain.SetDateActualFrom(dateActualFrom)
	} else {
		logger.Error("Error parsing DateActualFrom", "error", err)
	}

	if dateActualTo, err := time.Parse("2006-01-02T15:04:05Z", updatedRemainJSON.DateActualTo); err == nil {
		newRemain.SetDateActualTo(dateActualTo)
	} else {
		logger.Error("Error parsing DateActualTo", "error", err)
	}

	debugging.Remains = append(debugging.Remains, newRemain)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":      "Remain updated successfully",
		"index_remain": idRemain,
		"old_remain":   oldRemainJSON,
		"new_remain":   updatedRemainJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated remain", "status", http.StatusOK)
}

// delete
func RemainDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("DeleteRemain called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Error("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/remain/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idRemain, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldRemain *models.Remain
	for i, remain := range debugging.Remains {
		if remain.GetIdRemains() == idRemain {
			oldRemain = remain

			debugging.Remains = append(debugging.Remains[:i], debugging.Remains[i+1:]...)
			break
		}
	}

	if oldRemain == nil {
		http.Error(w, u.JsonErrorResponse("Remain not found"), http.StatusNotFound)
		return
	}

	oldRemainJSON, err := oldRemain.ToJSON()
	if err != nil {
		logger.Error("Error converting old remain to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old remain"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":      "Remain deleted successfully",
		"index_remain": idRemain,
		"old_remain":   oldRemainJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted remain", "status", http.StatusOK)
}
