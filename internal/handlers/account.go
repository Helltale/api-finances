package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/debuging"
	"github.com/helltale/api-finances/internal/models"
	u "github.com/helltale/api-finances/internal/utils"
)

// get all
func GetAllAccounts(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAllAccounts called", "method", r.Method)
	loggerFile.Info("GetAllAccounts called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var accounts []*models.Account
	if config.Mode == "debug" {
		accounts = debuging.Accounts
	} else {
		accounts = []*models.Account{}
	}

	var accountsJSON []models.AccountJSON
	for _, account := range accounts {
		accountJSON, err := account.ToJSON()
		if err != nil {
			loggerConsole.Error("Error converting account to JSON", "error", err)
			loggerFile.Error("Error converting account to JSON", "error", err)

			http.Error(w, u.JsonErrorResponse("Error converting account to JSON"), http.StatusInternalServerError)
			return
		}
		accountsJSON = append(accountsJSON, *accountJSON)
	}

	if err := json.NewEncoder(w).Encode(accountsJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved accounts", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved accounts", "status", http.StatusOK)
}

// get one by id
func GetAccountById(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("GetAccountById called", "method", r.Method)
	loggerFile.Info("GetAccountById called", "method", r.Method)

	if r.Method != http.MethodGet {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
	if config.Mode == "debug" {
		for _, account := range debuging.Accounts {
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
		loggerConsole.Error("Error converting account to JSON", "error", err)
		loggerFile.Error("Error converting account to JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error converting account to JSON"), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountJSON); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully retrieved account", "status", http.StatusOK)
	loggerFile.Info("Successfully retrieved account", "status", http.StatusOK)
}

// create
func PostAccount(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PostAccount called", "method", r.Method)
	loggerFile.Info("PostAccount called", "method", r.Method)

	if r.Method != http.MethodPost {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

		http.Error(w, u.JsonErrorResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var newAccountJSON models.AccountJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newAccountJSON); err != nil {
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	newAccount := &models.Account{}
	newAccount.SetIdAccaunt(newAccountJSON.IdAccaunt)
	newAccount.SetTgId(newAccountJSON.TgId)
	newAccount.SetName(newAccountJSON.Name)
	newAccount.SetGroupId(newAccountJSON.GroupId)

	debuging.Accounts = append(debuging.Accounts, newAccount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Account created successfully",
		"account": newAccountJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully created account", "status", http.StatusCreated)
	loggerFile.Info("Successfully created account", "status", http.StatusCreated)
}

// update
func PutAccount(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("PutAccount called", "method", r.Method)
	loggerFile.Info("PutAccount called", "method", r.Method)

	if r.Method != http.MethodPut {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
		loggerConsole.Error("Error decoding JSON", "error", err)
		loggerFile.Error("Error decoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Invalid JSON"), http.StatusBadRequest)
		return
	}

	var oldAccount *models.Account
	for i, account := range debuging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			oldAccount = account

			debuging.Accounts = append(debuging.Accounts[:i], debuging.Accounts[i+1:]...)
			break
		}
	}

	if oldAccount == nil {
		http.Error(w, u.JsonErrorResponse("Account not found"), http.StatusNotFound)
		return
	}

	oldAccountJSON, err := oldAccount.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old account to JSON", "error", err)
		http.Error(w, u.JsonErrorResponse("Error processing old account"), http.StatusInternalServerError)
		return
	}

	newAccount := &models.Account{}
	newAccount.SetIdAccaunt(updatedAccountJSON.IdAccaunt)
	newAccount.SetTgId(updatedAccountJSON.TgId)
	newAccount.SetName(updatedAccountJSON.Name)
	newAccount.SetGroupId(updatedAccountJSON.GroupId)

	debuging.Accounts = append(debuging.Accounts, newAccount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"message":       "Account updated successfully",
		"index_account": idAccaunt,
		"old_account":   oldAccountJSON,
		"new_account":   updatedAccountJSON,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully updated account", "status", http.StatusOK)
	loggerFile.Info("Successfully updated account", "status", http.StatusOK)
}

// delete
func DeleteAccount(w http.ResponseWriter, r *http.Request, loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	loggerConsole.Info("DeleteAccount called", "method", r.Method)
	loggerFile.Info("DeleteAccount called", "method", r.Method)

	if r.Method != http.MethodDelete {
		loggerConsole.Warn("Method not allowed", "method", r.Method)
		loggerFile.Warn("Method not allowed", "method", r.Method)

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
	for i, account := range debuging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			oldAccount = account

			debuging.Accounts = append(debuging.Accounts[:i], debuging.Accounts[i+1:]...)
			break
		}
	}

	if oldAccount == nil {
		http.Error(w, u.JsonErrorResponse("Account not found"), http.StatusNotFound)
		return
	}

	oldAccountJSON, err := oldAccount.ToJSON()
	if err != nil {
		loggerConsole.Error("Error converting old account to JSON", "error", err)
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
		loggerConsole.Error("Error encoding JSON", "error", err)
		loggerFile.Error("Error encoding JSON", "error", err)

		http.Error(w, u.JsonErrorResponse("Error encoding JSON"), http.StatusInternalServerError)
		return
	}

	loggerConsole.Info("Successfully deleted account", "status", http.StatusOK)
	loggerFile.Info("Successfully deleted account", "status", http.StatusOK)
}
