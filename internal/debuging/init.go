package debuging

import (
	"time"

	"github.com/helltale/api-finances/internal/models"
)

var (
	Incomes         []models.Income
	IncomesExpected []models.IncomeExpected
	Accounts        []models.Account
	Expences        []models.Expence
)

func Init() {
	income()
	incomeExpected()
	account()
	expence()
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

	Incomes = []models.Income{income1, income2}
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

	IncomesExpected = []models.IncomeExpected{incomeExpected1, incomeExpected2}
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

	Accounts = []models.Account{account1, account2}
}

func expence() {
	expence1 := models.Expence{}
	expence1.SetIdExpence(1)
	expence1.SetGroupExpence("Utilities")
	expence1.SetTitleExpence("Electricity Bill")
	expence1.SetDescriptionExpence("Monthly electricity bill payment")
	expence1.SetRepeat(1) // 1 - ежемесячно
	expence1.SetAmount(100.0)
	expence1.SetDate(time.Now())
	expence1.SetUpdBy("admin")
	expence1.SetDateActualFrom(time.Now())
	expence1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	expence2 := models.Expence{}
	expence2.SetIdExpence(2)
	expence2.SetGroupExpence("Groceries")
	expence2.SetTitleExpence("Weekly Groceries")
	expence2.SetDescriptionExpence("Weekly grocery shopping")
	expence2.SetRepeat(0) // 0 - единоразовая
	expence2.SetAmount(50.0)
	expence2.SetDate(time.Now())
	expence2.SetUpdBy("admin")
	expence2.SetDateActualFrom(time.Now())
	expence2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	Expences = []models.Expence{expence1, expence2}
}
