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
