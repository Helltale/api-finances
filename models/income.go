package models

type Income struct {
	id     int
	amount float64
	source string
}

// for serialization
type IncomeJSON struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Source string  `json:"source"`
}

func (i *Income) GetID() int {
	return i.id
}

func (i *Income) SetID(id int) {
	i.id = id
}

func (i *Income) GetAmount() float64 {
	return i.amount
}

func (i *Income) SetAmount(amount float64) {
	i.amount = amount
}

func (i *Income) GetSource() string {
	return i.source
}

func (i *Income) SetSource(source string) {
	i.source = source
}

func (i *Income) ToJSON() IncomeJSON {
	return IncomeJSON{
		ID:     i.id,
		Amount: i.amount,
		Source: i.source,
	}
}
