package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/logger"
	"github.com/helltale/api-finances/internal/models"
	"github.com/helltale/api-finances/internal/services"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func AccountGetAll(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAllAccounts called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var accounts []*models.Account
	if config.AppMode == "debug" {
		accounts = debugging.Accounts
	} else {
		accounts = []*models.Account{}
	}

	var accountsJSON []models.AccountJSON
	for _, account := range accounts {
		accountJSON, err := account.ToJSON()
		if err != nil {
			logger.Error("Error converting account to JSON", "error", err)
			http.Error(w, u.JsonErrorResponse("Error converting account to JSON"), http.StatusInternalServerError)
			return
		}
		accountsJSON = append(accountsJSON, *accountJSON)
	}

	if err := json.NewEncoder(w).Encode(accountsJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved accounts", "status", http.StatusOK)
}

// get one by id
func AccountGetByIdAccount(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("GetAccountById called", "method", r.Method)

	if r.Method != http.MethodGet {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idAccauntStr := strings.TrimPrefix(r.URL.Path, "/account/id/")
	if idAccauntStr == "" {
		http.Error(w, u.JsonErrorResponse("id_accaunt is required"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(idAccauntStr, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid id_accaunt"), http.StatusBadRequest)
		return
	}

	var foundAccount *models.Account
	if config.AppMode == "debug" {
		for _, account := range debugging.Accounts {
			if account.GetIdAccaunt() == idAccaunt {
				foundAccount = account
				break
			}
		}
	}

	if foundAccount == nil {
		http.Error(w, u.JsonErrorResponse("Account not found"), http.StatusNotFound)
		return
	}

	accountJSON, err := foundAccount.ToJSON()
	if err != nil {
		logger.Error("Error converting account to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error converting account to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountJSON); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved account", "status", http.StatusOK)
}

// create
func AccountPost(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PostAccount called", "method", r.Method)

	if r.Method != http.MethodPost {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newAccountJSON models.AccountJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newAccountJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newAccount := &models.Account{}
	newAccount.SetIdAccaunt(newAccountJSON.IdAccaunt)
	newAccount.SetTgId(newAccountJSON.TgId)
	newAccount.SetName(newAccountJSON.Name)
	newAccount.SetGroupId(newAccountJSON.GroupId)

	// service
	accountService := services.NewAccountService()
	if err := accountService.AddNewAccount(newAccount); err != nil {
		logger.Error("error adding account", "error", err)
		http.Error(w, u.JsonErrorResponse(err.Error()), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Account created successfully",
		"account": newAccountJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created account", "status", http.StatusCreated)
}

// update
func AccountPut(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("PutAccount called", "method", r.Method)

	if r.Method != http.MethodPut {
		logger.Info("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/account/update/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var updatedAccountJSON models.AccountJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedAccountJSON); err != nil {
		logger.Error("Error decoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldAccount *models.Account
	for i, account := range debugging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			oldAccount = account

			debugging.Accounts = append(debugging.Accounts[:i], debugging.Accounts[i+1:]...)
			break
		}
	}

	if oldAccount == nil {
		http.Error(w, u.JsonErrorResponse("Account not found"), http.StatusNotFound)
		return
	}

	oldAccountJSON, err := oldAccount.ToJSON()
	if err != nil {
		logger.Error("Error converting old account to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old account"), http.StatusInternalServerError)
		return
	}

	newAccount := &models.Account{}
	newAccount.SetIdAccaunt(updatedAccountJSON.IdAccaunt)
	newAccount.SetTgId(updatedAccountJSON.TgId)
	newAccount.SetName(updatedAccountJSON.Name)
	newAccount.SetGroupId(updatedAccountJSON.GroupId)

	debugging.Accounts = append(debugging.Accounts, newAccount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":       "Account updated successfully",
		"index_account": idAccaunt,
		"old_account":   oldAccountJSON,
		"new_account":   updatedAccountJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated account", "status", http.StatusOK)
}

// delete
func AccountDelete(w http.ResponseWriter, r *http.Request, logger *logger.CombinedLogger, config *config.Config) {
	logger.Info("DeleteAccount called", "method", r.Method)

	if r.Method != http.MethodDelete {
		logger.Info("Method not allowed", "method", r.Method)
		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	index := strings.TrimPrefix(urlPath, "/account/delete/")
	if index == urlPath {
		http.Error(w, u.JsonErrorResponse("Invalid URL"), http.StatusBadRequest)
		return
	}

	idAccaunt, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		http.Error(w, u.JsonErrorResponse("Invalid index format"), http.StatusBadRequest)
		return
	}

	var oldAccount *models.Account
	for i, account := range debugging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			oldAccount = account

			debugging.Accounts = append(debugging.Accounts[:i], debugging.Accounts[i+1:]...)
			break
		}
	}

	if oldAccount == nil {
		http.Error(w, u.JsonErrorResponse("Account not found"), http.StatusNotFound)
		return
	}

	oldAccountJSON, err := oldAccount.ToJSON()
	if err != nil {
		logger.Error("Error converting old account to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old account"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":       "Account deleted successfully",
		"index_account": idAccaunt,
		"old_account":   oldAccountJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted account", "status", http.StatusOK)
}
