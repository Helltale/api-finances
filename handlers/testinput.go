package handlers

import (
	"github.com/helltale/api-finances/models"
)

func Init() {
	income1 := models.Income{}
	income1.SetID(1)
	income1.SetAmount(100.5)
	income1.SetSource("Salary")

	income2 := models.Income{}
	income2.SetID(2)
	income2.SetAmount(50.5)
	income2.SetSource("Freelance")

	incomes = []models.Income{income1, income2}
}
