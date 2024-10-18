package models

import "time"

type Cashback struct {
	idCashback int64
	idAccaunt  int64
	bankName   string
	category   string
	percent    int8

	updBy          string    // who changed
	dateActualFrom time.Time // actual from
	dateActualTo   time.Time // actual to if 9999-12-31 to now
}

type CashbackJSON struct {
	IdCashback     int64  `json:"id_cashback"`
	IdAccaunt      int64  `json:"id_accaunt"`
	BankName       string `json:"bank_name"`
	Category       string `json:"category"`
	Percent        int8   `json:"percent"`
	UpdBy          string `json:"upd_by"`
	DateActualFrom string `json:"date_actual_from"`
	DateActualTo   string `json:"date_actual_to"`
}

func (c *Cashback) ToJSON() (*CashbackJSON, error) {
	return &CashbackJSON{
		IdCashback:     c.idCashback,
		IdAccaunt:      c.idAccaunt,
		BankName:       c.bankName,
		Category:       c.category,
		Percent:        c.percent,
		UpdBy:          c.updBy,
		DateActualFrom: c.dateActualFrom.Format("2006-01-02 15:04:05"),
		DateActualTo:   c.dateActualTo.Format("2006-01-02 15:04:05"),
	}, nil
}

func (c *Cashback) GetIdCashback() int64 {
	return c.idCashback
}

func (c *Cashback) GetIdAccaunt() int64 {
	return c.idAccaunt
}

func (c *Cashback) GetBankName() string {
	return c.bankName
}

func (c *Cashback) GetCategory() string {
	return c.category
}

func (c *Cashback) GetPercent() int8 {
	return c.percent
}

func (c *Cashback) GetUpdBy() string {
	return c.updBy
}

func (c *Cashback) GetDateActualFrom() time.Time {
	return c.dateActualFrom
}

func (c *Cashback) GetDateActualTo() time.Time {
	return c.dateActualTo
}

func (c *Cashback) SetIdCashback(id int64) {
	c.idCashback = id
}

func (c *Cashback) SetIdAccaunt(id int64) {
	c.idAccaunt = id
}

func (c *Cashback) SetBankName(name string) {
	c.bankName = name
}

func (c *Cashback) SetCategory(category string) {
	c.category = category
}

func (c *Cashback) SetPercent(percent int8) {
	c.percent = percent
}

func (c *Cashback) SetUpdBy(updBy string) {
	c.updBy = updBy
}

func (c *Cashback) SetDateActualFrom(date time.Time) {
	c.dateActualFrom = date
}

func (c *Cashback) SetDateActualTo(date time.Time) {
	c.dateActualTo = date
}
