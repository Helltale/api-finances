package models

import (
	"time"
)

type IncomeExpected struct {
	id_accaunt       int64
	id_income_ex     int64
	amount           float64
	type_income      string //salary or award
	income_mont_date int    //1-31

	u_by             string    //who changed
	date_actual_from time.Time //actual from
	date_actual_to   time.Time //actual to if 9999-12-31 to nowday
}
