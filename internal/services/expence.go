package services

import (
	"errors"
	"time"

	"github.com/helltale/api-finances/internal/debugging"
	"github.com/helltale/api-finances/internal/models"
)

type ExpenceService struct{}

func NewExpenceService() *ExpenceService {
	return &ExpenceService{}
}

func (s *ExpenceService) AddNewExpence(newExpence *models.Expence) error {
	for _, expence := range debugging.Expences {
		if expence.GetIdExpence() == newExpence.GetIdExpence() {
			return errors.New("expence with this ID already exists")
		}
	}
	debugging.Expences = append(debugging.Expences, newExpence)
	return nil
}

func (s *ExpenceService) GetAllExpences() []*models.Expence {
	return debugging.Expences
}

func (s *ExpenceService) GetExpenceById(idExpence int64) (*models.Expence, error) {
	for _, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			return expence, nil
		}
	}
	return nil, errors.New("expence not found")
}

func (s *ExpenceService) UpdateExpence(updatedExpence *models.Expence) (*models.Expence, error) {
	for i, expence := range debugging.Expences {
		if expence.GetIdExpence() == updatedExpence.GetIdExpence() {
			oldExpenceCopy := &models.Expence{}
			*oldExpenceCopy = *expence

			debugging.Expences[i] = updatedExpence
			return oldExpenceCopy, nil
		}
	}
	return nil, errors.New("expence not found")
}

func (s *ExpenceService) UpdateHistoryExpence(idExpence int64, newExpence *models.Expence) (*models.Expence, error) {
	today := time.Now()

	var oldExpence *models.Expence
	for i, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			oldExpence = expence
			debugging.Expences[i].SetDateActualTo(today)
			break
		}
	}

	if oldExpence == nil {
		return nil, errors.New("expence not found")
	}

	newExpence.SetDateActualFrom(today)
	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	newExpence.SetDateActualTo(futureDate)

	debugging.Expences = append(debugging.Expences, newExpence)

	return oldExpence, nil
}

func (s *ExpenceService) DeleteExpence(idExpence int64) (*models.Expence, error) {
	for i, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			debugging.Expences = append(debugging.Expences[:i], debugging.Expences[i+1:]...)
			return expence, nil
		}
	}
	return nil, errors.New("expence not found")
}

func (service *ExpenceService) GetExpencesByGroup(group string) ([]*models.Expence, error) {
	var expences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetGroupExpence() == group {
			expences = append(expences, expence)
		}
	}

	if len(expences) == 0 {
		return nil, errors.New("no expences found for the specified group")
	}

	return expences, nil
}

func (s *ExpenceService) DeleteAndRestorePreviousExpence(idExpence int64) (*models.Expence, error) {
	var currentRecord *models.Expence
	var lastHistoricalRecord *models.Expence
	maxDate := time.Time{}

	for _, expence := range debugging.Expences {
		if expence.GetIdExpence() == idExpence {
			if expence.GetDateActualTo().Year() == 9999 {
				currentRecord = expence
			} else if expence.GetDateActualTo().After(maxDate) {
				lastHistoricalRecord = expence
				maxDate = expence.GetDateActualTo()
			}
		}
	}

	if currentRecord == nil {
		return nil, errors.New("active expence record not found")
	}

	if lastHistoricalRecord == nil {
		return nil, errors.New("no historical record found to restore")
	}

	for i, expence := range debugging.Expences {
		if expence == currentRecord {
			debugging.Expences = append(debugging.Expences[:i], debugging.Expences[i+1:]...)
			break
		}
	}

	futureDate, _ := time.Parse("2006-01-02", "9999-12-31")
	lastHistoricalRecord.SetDateActualTo(futureDate)

	return lastHistoricalRecord, nil
}

func (service *ExpenceService) GetExpencesByTitle(title string) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetTitleExpence() == title {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found for the specified title")
	}

	return foundExpences, nil
}

func (service *ExpenceService) GetExpencesByDateRange(startDate, endDate time.Time) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetDate().After(startDate) && expence.GetDate().Before(endDate) {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found in the specified date range")
	}

	return foundExpences, nil
}

func (service *ExpenceService) GetExpencesByRepeat(repeat int8) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetRepeat() == repeat {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found for the specified repeat type")
	}

	return foundExpences, nil
}

func (service *ExpenceService) GetExpencesByAmountRange(minAmount, maxAmount float64) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetAmount() >= minAmount && expence.GetAmount() <= maxAmount {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found in the specified amount range")
	}

	return foundExpences, nil
}

func (service *ExpenceService) GetExpencesByMaxAmount(maxAmount float64) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetAmount() < maxAmount {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found below the specified amount")
	}

	return foundExpences, nil
}

func (service *ExpenceService) GetExpencesByMinAmount(minAmount float64) ([]*models.Expence, error) {
	var foundExpences []*models.Expence

	for _, expence := range debugging.Expences {
		if expence.GetAmount() > minAmount {
			foundExpences = append(foundExpences, expence)
		}
	}

	if len(foundExpences) == 0 {
		return nil, errors.New("no expences found above the specified amount")
	}

	return foundExpences, nil
}
