package models

import "time"

type Remain struct {
	idRemains        int64     // id
	idAccaunt        int64     // account id
	amount           float64   // actual sum
	lastUpdateAmount float64   // sum of last operation
	lastUpdateId     int64     // id of last operation
	lastUpdateGroup  string    // name of last operation
	updBy            string    // who changed
	dateActualFrom   time.Time // actual from
	dateActualTo     time.Time // actual to if 9999-12-31 to now
}

type RemainJSON struct {
	IdRemains        int64   `json:"id_remains"`
	IdAccaunt        int64   `json:"id_accaunt"`
	Amount           float64 `json:"amount"`
	LastUpdateAmount float64 `json:"last_update_amount"`
	LastUpdateId     int64   `json:"last_update_id"`
	LastUpdateGroup  string  `json:"last_update_group"`
	UpdBy            string  `json:"upd_by"`
	DateActualFrom   string  `json:"date_actual_from"`
	DateActualTo     string  `json:"date_actual_to"`
}

func (r *Remain) ToJSON() (*RemainJSON, error) {
	return &RemainJSON{
		IdRemains:        r.idRemains,
		IdAccaunt:        r.idAccaunt,
		Amount:           r.amount,
		LastUpdateAmount: r.lastUpdateAmount,
		LastUpdateId:     r.lastUpdateId,
		LastUpdateGroup:  r.lastUpdateGroup,
		UpdBy:            r.updBy,
		DateActualFrom:   r.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:     r.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *Remain) GetIdRemains() int64 {
	return r.idRemains
}

func (r *Remain) GetIdAccaunt() int64 {
	return r.idAccaunt
}

func (r *Remain) GetAmount() float64 {
	return r.amount
}

func (r *Remain) GetLastUpdateAmount() float64 {
	return r.lastUpdateAmount
}

func (r *Remain) GetLastUpdateId() int64 {
	return r.lastUpdateId
}

func (r *Remain) GetLastUpdateGroup() string {
	return r.lastUpdateGroup
}

func (r *Remain) GetUpdBy() string {
	return r.updBy
}

func (r *Remain) GetDateActualFrom() time.Time {
	return r.dateActualFrom
}

func (r *Remain) GetDateActualTo() time.Time {
	return r.dateActualTo
}

func (r *Remain) SetIdRemains(id int64) {
	r.idRemains = id
}

func (r *Remain) SetIdAccaunt(id int64) {
	r.idAccaunt = id
}

func (r *Remain) SetAmount(amount float64) {
	r.amount = amount
}

func (r *Remain) SetLastUpdateAmount(lastUpdateAmount float64) {
	r.lastUpdateAmount = lastUpdateAmount
}

func (r *Remain) SetLastUpdateId(lastUpdateId int64) {
	r.lastUpdateId = lastUpdateId
}

func (r *Remain) SetLastUpdateGroup(lastUpdateGroup string) {
	r.lastUpdateGroup = lastUpdateGroup
}

func (r *Remain) SetUpdBy(updBy string) {
	r.updBy = updBy
}

func (r *Remain) SetDateActualFrom(date time.Time) {
	r.dateActualFrom = date
}

func (r *Remain) SetDateActualTo(date time.Time) {
	r.dateActualTo = date
}
