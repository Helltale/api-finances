package models

import (
	"time"
)

type IncomeExpected struct {
	idAccaunt       int64
	idIncomeEx      int64
	amount          float64
	typeIncome      string    //salary or award
	incomeMonthDate int8      //1-31
	updBy           string    //who changed
	dateActualFrom  time.Time //actual from
	dateActualTo    time.Time //actual to if 9999-12-31 to nowday
}

type IncomeExpectedJSON struct {
	IdAccaunt       int64   `json:"id_accaunt"`
	IdIncomeEx      int64   `json:"id_income_ex"`
	Amount          float64 `json:"amount"`
	TypeIncome      string  `json:"type_income"`
	IncomeMonthDate int8    `json:"income_month_date"`
	UpdBy           string  `json:"upd_by"`
	DateActualFrom  string  `json:"date_actual_from"`
	DateActualTo    string  `json:"date_actual_to"`
}

func (ie *IncomeExpected) ToJSON() (*IncomeExpectedJSON, error) {
	return &IncomeExpectedJSON{
		IdAccaunt:       ie.idAccaunt,
		IdIncomeEx:      ie.idIncomeEx,
		Amount:          ie.amount,
		TypeIncome:      ie.typeIncome,
		IncomeMonthDate: ie.incomeMonthDate,
		UpdBy:           ie.updBy,
		DateActualFrom:  ie.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:    ie.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (ie *IncomeExpected) GetIdAccaunt() int64 {
	return ie.idAccaunt
}

func (ie *IncomeExpected) GetIdIncomeEx() int64 {
	return ie.idIncomeEx
}

func (ie *IncomeExpected) GetAmount() float64 {
	return ie.amount
}

func (ie *IncomeExpected) GetTypeIncome() string {
	return ie.typeIncome
}

func (ie *IncomeExpected) GetIncomeMonthDate() int8 {
	return ie.incomeMonthDate
}

func (ie *IncomeExpected) GetUpdBy() string {
	return ie.updBy
}

func (ie *IncomeExpected) GetDateActualFrom() time.Time {
	return ie.dateActualFrom
}

func (ie *IncomeExpected) GetDateActualTo() time.Time {
	return ie.dateActualTo
}

func (ie *IncomeExpected) SetIdAccaunt(id int64) {
	ie.idAccaunt = id
}

func (ie *IncomeExpected) SetIdIncomeEx(id int64) {
	ie.idIncomeEx = id
}

func (ie *IncomeExpected) SetAmount(amount float64) {
	ie.amount = amount
}

func (ie *IncomeExpected) SetTypeIncome(typeIncome string) {
	ie.typeIncome = typeIncome
}

func (ie *IncomeExpected) SetIncomeMonthDate(date int8) {
	ie.incomeMonthDate = date
}

func (ie *IncomeExpected) SetUpdBy(updBy string) {
	ie.updBy = updBy
}

func (ie *IncomeExpected) SetDateActualFrom(date time.Time) {
	ie.dateActualFrom = date
}

func (ie *IncomeExpected) SetDateActualTo(date time.Time) {
	ie.dateActualTo = date
}
