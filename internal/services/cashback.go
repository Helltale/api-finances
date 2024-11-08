package services

import (
	"errors"
	"time"

	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/models"
)

type CashbackService struct{}

func NewCashbackService() *CashbackService {
	return &CashbackService{}
}

func (s *CashbackService) AddNewCashback(newCashback *models.Cashback) error {
	for _, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == newCashback.GetIdCashback() {
			return errors.New("cashback with this ID already exists")
		}
	}

	debugging.Cashbacks = append(debugging.Cashbacks, newCashback)
	return nil
}

func (s *CashbackService) GetAllCashbacks() []*models.Cashback {
	return debugging.Cashbacks
}

func (s *CashbackService) GetCashbackById(idCashback int64) (*models.Cashback, error) {
	for _, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			return cashback, nil
		}
	}
	return nil, errors.New("cashback not found")
}

func (s *CashbackService) UpdateCashback(updatedCashback *models.Cashback) (*models.Cashback, error) {
	for i, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == updatedCashback.GetIdCashback() {
			oldCashbackCopy := &models.Cashback{}
			*oldCashbackCopy = *cashback

			debugging.Cashbacks[i] = updatedCashback
			return oldCashbackCopy, nil
		}
	}
	return nil, errors.New("cashback not found")
}

func (s *CashbackService) UpdateHistoryCashback(idCashback int64, newCashback *models.Cashback) (*models.Cashback, error) {
	today := time.Now()

	var oldCashback *models.Cashback
	for i, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			oldCashback = cashback
			debugging.Cashbacks[i].SetDateActualTo(today)
			break
		}
	}

	if oldCashback == nil {
		return nil, errors.New("cashback not found")
	}

	newCashback.SetDateActualFrom(today)
	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	newCashback.SetDateActualTo(futureDate)

	debugging.Cashbacks = append(debugging.Cashbacks, newCashback)

	return oldCashback, nil
}

func (s *CashbackService) DeleteCashback(idCashback int64) (*models.Cashback, error) {
	for i, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			debugging.Cashbacks = append(debugging.Cashbacks[:i], debugging.Cashbacks[i+1:]...)
			return cashback, nil
		}
	}
	return nil, errors.New("cashback not found")
}

func (s *CashbackService) DeleteAndRestorePreviousCashback(idCashback int64) (*models.Cashback, error) {
	var currentRecord *models.Cashback
	var lastHistoricalRecord *models.Cashback
	maxDate := time.Time{}

	for _, cashback := range debugging.Cashbacks {
		if cashback.GetIdCashback() == idCashback {
			if cashback.GetDateActualTo().Year() == 9999 {
				currentRecord = cashback
			} else if cashback.GetDateActualTo().After(maxDate) {
				lastHistoricalRecord = cashback
				maxDate = cashback.GetDateActualTo()
			}
		}
	}

	if currentRecord == nil {
		return nil, errors.New("active cashback record not found")
	}

	if lastHistoricalRecord == nil {
		return nil, errors.New("no historical record found to restore")
	}

	for i, cashback := range debugging.Cashbacks {
		if cashback == currentRecord {
			debugging.Cashbacks = append(debugging.Cashbacks[:i], debugging.Cashbacks[i+1:]...)
			break
		}
	}

	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	lastHistoricalRecord.SetDateActualTo(futureDate)

	return lastHistoricalRecord, nil
}
