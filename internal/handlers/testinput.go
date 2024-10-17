package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/models"
)

func Init(config config.Config) {
	switch config.Mode {
	case "debug":
		income()
		incomeExpected()
		account()
		fmt.Println("info: run in debug mode")
	case "release":
		fmt.Println("info: run release mode")
	default:
		log.Panic("bad config.mode")
	}

}

func income() {
	income1 := models.Income{}
	income1.SetIdIncome(1)
	income1.SetAmount(100.5)
	income1.SetTypeIncome("Salary")
	income1.SetIdAccaunt(1)
	income1.SetIdIncomeExpected(1)
	income1.SetIncomeMonthMonth(1)
	income1.SetIncomeMonthDate(15)
	income1.SetUpdBy("admin")
	income1.SetDateActualFrom(time.Now())
	income1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	income2 := models.Income{}
	income2.SetIdIncome(2)
	income2.SetAmount(50.5)
	income2.SetTypeIncome("Freelance")
	income2.SetIdAccaunt(2)
	income2.SetIdIncomeExpected(2)
	income2.SetIncomeMonthMonth(1)
	income2.SetIncomeMonthDate(20)
	income2.SetUpdBy("admin")
	income2.SetDateActualFrom(time.Now())
	income2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	incomes = []models.Income{income1, income2}
}

func incomeExpected() {
	incomeExpected1 := models.IncomeExpected{}
	incomeExpected1.SetIdIncomeEx(1)
	incomeExpected1.SetAmount(150.0)
	incomeExpected1.SetTypeIncome("Salary")
	incomeExpected1.SetIdAccaunt(1)
	incomeExpected1.SetIncomeMonthDate(15)
	incomeExpected1.SetUpdBy("admin")
	incomeExpected1.SetDateActualFrom(time.Now())
	incomeExpected1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	incomeExpected2 := models.IncomeExpected{}
	incomeExpected2.SetIdIncomeEx(2)
	incomeExpected2.SetAmount(75.0)
	incomeExpected2.SetTypeIncome("Bonus")
	incomeExpected2.SetIdAccaunt(2)
	incomeExpected2.SetIncomeMonthDate(20)
	incomeExpected2.SetUpdBy("admin")
	incomeExpected2.SetDateActualFrom(time.Now())
	incomeExpected2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	incomesExpected = []models.IncomeExpected{incomeExpected1, incomeExpected2}
}

func account() {
	account1 := models.Account{}
	account1.SetIdAccaunt(1)
	account1.SetTgId(123456789)
	account1.SetName("user1")
	account1.SetGroupId(1)

	account2 := models.Account{}
	account2.SetIdAccaunt(2)
	account2.SetTgId(987654321)
	account2.SetName("user2")
	account2.SetGroupId(1)

	accounts = []models.Account{account1, account2}
}
