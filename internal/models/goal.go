package models

import "time"

type Goal struct {
	idGoal         int64     // id
	idAccaunt      int64     // account id
	amount         float64   // actual sum
	date           time.Time // deadline
	updBy          string    // who changed
	dateActualFrom time.Time // actual from
	dateActualTo   time.Time // actual to if 9999-12-31 to now
}

type GoalJSON struct {
	IdGoal         int64   `json:"id_goal"`
	IdAccaunt      int64   `json:"id_accaunt"`
	Amount         float64 `json:"amount"`
	Date           string  `json:"date"`
	UpdBy          string  `json:"upd_by"`
	DateActualFrom string  `json:"date_actual_from"`
	DateActualTo   string  `json:"date_actual_to"`
}

func (g *Goal) ToJSON() (*GoalJSON, error) {
	return &GoalJSON{
		IdGoal:         g.idGoal,
		IdAccaunt:      g.idAccaunt,
		Amount:         g.amount,
		Date:           g.date.Format("2006-01-02 15:04:05"),
		UpdBy:          g.updBy,
		DateActualFrom: g.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:   g.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (g *Goal) GetIdGoal() int64 {
	return g.idGoal
}

func (g *Goal) GetIdAccaunt() int64 {
	return g.idAccaunt
}

func (g *Goal) GetAmount() float64 {
	return g.amount
}

func (g *Goal) GetDate() time.Time {
	return g.date
}

func (g *Goal) GetUpdBy() string {
	return g.updBy
}

func (g *Goal) GetDateActualFrom() time.Time {
	return g.dateActualFrom
}

func (g *Goal) GetDateActualTo() time.Time {
	return g.dateActualTo
}

func (g *Goal) SetIdGoal(id int64) {
	g.idGoal = id
}

func (g *Goal) SetIdAccaunt(id int64) {
	g.idAccaunt = id
}

func (g *Goal) SetAmount(amount float64) {
	g.amount = amount
}

func (g *Goal) SetDate(date time.Time) {
	g.date = date
}

func (g *Goal) SetUpdBy(updBy string) {
	g.updBy = updBy
}

func (g *Goal) SetDateActualFrom(date time.Time) {
	g.dateActualFrom = date
}

func (g *Goal) SetDateActualTo(date time.Time) {
	g.dateActualTo = date
}
