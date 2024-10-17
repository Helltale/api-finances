package models

import (
	"time"
)

// оставить одну модель на ежемес траты и единоразовые, но проверять через repeat + dateActualFrom + dateActualTo
type Expence struct {
	idExpence          int64  // айди траты
	groupExpence       string // группа траты
	titleExpence       string // название траты
	descriptionExpence string // доп инфа о трате
	repeat             int8   // ежемес или нет
	amount             float64
	date               time.Time // дата совершения единоразовой покупки
	updBy              string    // who changed
	dateActualFrom     time.Time // actual from
	dateActualTo       time.Time // actual to if 9999-12-31 to now
}

type ExpenceJSON struct {
	IdExpence          int64   `json:"id_expence"`
	GroupExpence       string  `json:"group_expence"`
	TitleExpence       string  `json:"title_expence"`
	DescriptionExpence string  `json:"description_expence"`
	Repeat             int8    `json:"repeat"`
	Amount             float64 `json:"amount"`
	Date               string  `json:"date"`
	UpdBy              string  `json:"upd_by"`
	DateActualFrom     string  `json:"date_actual_from"`
	DateActualTo       string  `json:"date_actual_to"`
}

func (e *Expence) ToJSON() (*ExpenceJSON, error) {
	return &ExpenceJSON{
		IdExpence:          e.idExpence,
		GroupExpence:       e.groupExpence,
		TitleExpence:       e.titleExpence,
		DescriptionExpence: e.descriptionExpence,
		Repeat:             e.repeat,
		Amount:             e.amount,
		Date:               e.date.Format("2006-01-02 15:04:05"),
		UpdBy:              e.updBy,
		DateActualFrom:     e.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:       e.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (e *Expence) GetIdExpence() int64 {
	return e.idExpence
}

func (e *Expence) GetGroupExpence() string {
	return e.groupExpence
}

func (e *Expence) GetTitleExpence() string {
	return e.titleExpence
}

func (e *Expence) GetDescriptionExpence() string {
	return e.descriptionExpence
}

func (e *Expence) GetRepeat() int8 {
	return e.repeat
}

func (e *Expence) GetAmount() float64 {
	return e.amount
}

func (e *Expence) GetDate() time.Time {
	return e.date
}

func (e *Expence) GetUpdBy() string {
	return e.updBy
}

func (e *Expence) GetDateActualFrom() time.Time {
	return e.dateActualFrom
}

func (e *Expence) GetDateActualTo() time.Time {
	return e.dateActualTo
}

func (e *Expence) SetIdExpence(id int64) {
	e.idExpence = id
}

func (e *Expence) SetGroupExpence(group string) {
	e.groupExpence = group
}

func (e *Expence) SetTitleExpence(title string) {
	e.titleExpence = title
}

func (e *Expence) SetDescriptionExpence(description string) {
	e.descriptionExpence = description
}

func (e *Expence) SetRepeat(repeat int8) {
	e.repeat = repeat
}

func (e *Expence) SetAmount(amount float64) {
	e.amount = amount
}

func (e *Expence) SetDate(date time.Time) {
	e.date = date
}

func (e *Expence) SetUpdBy(updBy string) {
	e.updBy = updBy
}

func (e *Expence) SetDateActualFrom(date time.Time) {
	e.dateActualFrom = date
}

func (e *Expence) SetDateActualTo(date time.Time) {
	e.dateActualTo = date
}
