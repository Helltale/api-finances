package models

import "time"

type CustomTime time.Time

type Income struct {
	idIncome         int64
	idAccaunt        int64
	idIncomeExpected int64   // id expected
	amount           float64 // real amount
	expectedAmount   float64 // expected amount
	typeIncome       string  // salary or award
	incomeMonthMonth int     // 1-12
	incomeMonthDate  int     // 1-31

	updBy          string    // who changed
	dateActualFrom time.Time // actual from
	dateActualTo   time.Time // actual to if 9999-12-31 to now
}

type IncomeJSON struct {
	IdIncome         int64   `json:"id_income"`
	IdAccaunt        int64   `json:"id_accaunt"`
	IdIncomeExpected int64   `json:"id_income_expected"`
	Amount           float64 `json:"amount"`
	ExpectedAmount   float64 `json:"expected_amount"`
	TypeIncome       string  `json:"type_income"`
	IncomeMonthMonth int     `json:"income_month_month"`
	IncomeMonthDate  int     `json:"income_month_date"`
	UpdBy            string  `json:"upd_by"`
	DateActualFrom   string  `json:"date_actual_from"`
	DateActualTo     string  `json:"date_actual_to"`
}

func (i *Income) ToJSON() (*IncomeJSON, error) {
	return &IncomeJSON{
		IdIncome:         i.idIncome,
		IdAccaunt:        i.idAccaunt,
		IdIncomeExpected: i.idIncomeExpected,
		Amount:           i.amount,
		ExpectedAmount:   i.expectedAmount,
		TypeIncome:       i.typeIncome,
		IncomeMonthMonth: i.incomeMonthMonth,
		IncomeMonthDate:  i.incomeMonthDate,
		UpdBy:            i.updBy,
		DateActualFrom:   i.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:     i.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (i *Income) GetIdIncome() int64 {
	return i.idIncome
}

func (i *Income) GetIdAccaunt() int64 {
	return i.idAccaunt
}

func (i *Income) GetIdIncomeExpected() int64 {
	return i.idIncomeExpected
}

func (i *Income) GetAmount() float64 {
	return i.amount
}

func (i *Income) GetExpectedAmount() float64 {
	return i.expectedAmount
}

func (i *Income) GetTypeIncome() string {
	return i.typeIncome
}

func (i *Income) GetIncomeMonthMonth() int {
	return i.incomeMonthMonth
}

func (i *Income) GetIncomeMonthDate() int {
	return i.incomeMonthDate
}

func (i *Income) GetUpdBy() string {
	return i.updBy
}

func (i *Income) GetDateActualFrom() time.Time {
	return i.dateActualFrom
}

func (i *Income) GetDateActualTo() time.Time {
	return i.dateActualTo
}

func (i *Income) SetIdIncome(id int64) {
	i.idIncome = id
}

func (i *Income) SetIdAccaunt(id int64) {
	i.idAccaunt = id
}

func (i *Income) SetIdIncomeExpected(id int64) {
	i.idIncomeExpected = id
}

func (i *Income) SetAmount(amount float64) {
	i.amount = amount
}

func (i *Income) SetExpectedAmount(expectedAmount float64) {
	i.expectedAmount = expectedAmount
}

func (i *Income) SetTypeIncome(typeIncome string) {
	i.typeIncome = typeIncome
}

func (i *Income) SetIncomeMonthMonth(month int) {
	i.incomeMonthMonth = month
}

func (i *Income) SetIncomeMonthDate(date int) {
	i.incomeMonthDate = date
}

func (i *Income) SetUpdBy(updBy string) {
	i.updBy = updBy
}

func (i *Income) SetDateActualFrom(date time.Time) {
	i.dateActualFrom = date
}

func (i *Income) SetDateActualTo(date time.Time) {
	i.dateActualTo = date
}
