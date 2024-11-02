package services

import (
	"errors"
	"time"

	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/models"
)

type IncomeExpectedService struct{}

func NewIncomeExpectedService() *IncomeExpectedService {
	return &IncomeExpectedService{}
}

func (s *IncomeExpectedService) AddNewIncomeExpected(newIncomeExpected *models.IncomeExpected) error {
	for _, incomeExpected := range debugging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == newIncomeExpected.GetIdIncomeEx() {
			return errors.New("income with this ID already exists")
		}
	}

	debugging.IncomesExpected = append(debugging.IncomesExpected, newIncomeExpected)
	return nil
}

func (s *IncomeExpectedService) GetAllIncomesExpected() []*models.IncomeExpected {
	return debugging.IncomesExpected
}

func (s *IncomeExpectedService) GetIncomeExpectedById(idIncomeEx int64) (*models.IncomeExpected, error) {
	for _, incomeExpected := range debugging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == idIncomeEx {
			return incomeExpected, nil
		}
	}
	return nil, errors.New("income expected not found")
}

func (s *IncomeExpectedService) UpdateIncomeExpected(updatedIncomeExpected *models.IncomeExpected) (*models.IncomeExpected, error) {
	for i, incomeExpected := range debugging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == updatedIncomeExpected.GetIdIncomeEx() {
			oldIncomeExpectedCopy := &models.IncomeExpected{}
			*oldIncomeExpectedCopy = *incomeExpected

			debugging.IncomesExpected[i] = updatedIncomeExpected
			return oldIncomeExpectedCopy, nil
		}
	}
	return nil, errors.New("income expected not found")
}

func (s *IncomeExpectedService) UpdateHistoryIncomeExpected(idIncomeEx int64, newIncomeExpected *models.IncomeExpected) (*models.IncomeExpected, error) {
	today := time.Now()

	var oldIncomeExpected *models.IncomeExpected
	for i, income := range debugging.IncomesExpected {
		if income.GetIdIncomeEx() == idIncomeEx {
			oldIncomeExpected = income
			debugging.IncomesExpected[i].SetDateActualTo(today)
			break
		}
	}

	if oldIncomeExpected == nil {
		return nil, errors.New("income expected not found")
	}

	newIncomeExpected.SetDateActualFrom(today)
	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	newIncomeExpected.SetDateActualTo(futureDate)

	debugging.IncomesExpected = append(debugging.IncomesExpected, newIncomeExpected)

	return oldIncomeExpected, nil
}

func (s *IncomeExpectedService) DeleteIncomeExpected(idIncomeEx int64) (*models.IncomeExpected, error) {
	for i, incomeExpected := range debugging.IncomesExpected {
		if incomeExpected.GetIdIncomeEx() == idIncomeEx {
			debugging.IncomesExpected = append(debugging.IncomesExpected[:i], debugging.IncomesExpected[i+1:]...)
			return debugging.IncomesExpected[i], nil
		}
	}
	return nil, errors.New("income expected not found")
}

func (s *IncomeExpectedService) DeleteAndRestorePreviousIncomeExpexted(idIncomeEx int64) (*models.IncomeExpected, error) {
	var currentRecord *models.IncomeExpected
	var lastHistoricalRecord *models.IncomeExpected
	maxDate := time.Time{}

	for _, income := range debugging.IncomesExpected {
		if income.GetIdIncomeEx() == idIncomeEx {
			if income.GetDateActualTo().Year() == 9999 {
				currentRecord = income
			} else if income.GetDateActualTo().After(maxDate) {

				lastHistoricalRecord = income
				maxDate = income.GetDateActualTo()
			}
		}
	}

	if currentRecord == nil {
		return nil, errors.New("active income expected record not found")
	}

	if lastHistoricalRecord == nil {
		return nil, errors.New("no historical record found to restore")
	}

	for i, income := range debugging.IncomesExpected {
		if income == currentRecord {
			debugging.IncomesExpected = append(debugging.IncomesExpected[:i], debugging.IncomesExpected[i+1:]...)
			break
		}
	}

	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	lastHistoricalRecord.SetDateActualTo(futureDate)

	return lastHistoricalRecord, nil
}
