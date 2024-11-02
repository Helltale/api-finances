package services

import (
	"errors"

	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/models"
)

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) AddNewAccount(newAccount *models.Account) error {
	for _, account := range debugging.Accounts {
		if account == newAccount {
			return errors.New("this entry already exists")
		}

		if account.GetTgId() == newAccount.GetTgId() {
			return errors.New("this tgid already exists")
		}
	}

	debugging.Accounts = append(debugging.Accounts, newAccount)
	return nil
}

func (s *AccountService) GetAllAccounts() []*models.Account {
	return debugging.Accounts
}

func (s *AccountService) GetAccountById(idAccaunt int64) (*models.Account, error) {
	for _, account := range debugging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

func (s *AccountService) UpdateAccount(updatedAccount *models.Account) (*models.Account, error) {
	for i, account := range debugging.Accounts {
		if account.GetIdAccaunt() == updatedAccount.GetIdAccaunt() {
			oldAccountCopy := &models.Account{}
			*oldAccountCopy = *account

			debugging.Accounts[i] = updatedAccount
			return oldAccountCopy, nil
		}
	}
	return nil, errors.New("account not found")
}

func (s *AccountService) DeleteAccount(idAccaunt int64) error {
	for i, account := range debugging.Accounts {
		if account.GetIdAccaunt() == idAccaunt {
			debugging.Accounts = append(debugging.Accounts[:i], debugging.Accounts[i+1:]...)
			return nil
		}
	}
	return errors.New("account not found")
}
