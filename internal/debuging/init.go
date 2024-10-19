package debuging

import (
	"time"

	"github.com/helltale/api-finances/internal/models"
)

var (
	Incomes         []*models.Income
	IncomesExpected []*models.IncomeExpected
	Accounts        []models.Account
	Expences        []models.Expence
	Remains         []models.Remain
	Goals           []models.Goal
	Cashbacks       []models.Cashback
)

func Init() {
	income()
	incomeExpected()
	account()
	expence()
	remain()
	goal()
	cashback()
}

func remain() {
	remain1 := models.Remain{}
	remain1.SetIdRemains(1)
	remain1.SetIdAccaunt(1)
	remain1.SetAmount(1500.75)
	remain1.SetLastUpdateAmount(200.00)
	remain1.SetLastUpdateId(1)
	remain1.SetLastUpdateGroup("Deposit")
	remain1.SetUpdBy("admin")
	remain1.SetDateActualFrom(time.Now())
	remain1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	remain2 := models.Remain{}
	remain2.SetIdRemains(2)
	remain2.SetIdAccaunt(2)
	remain2.SetAmount(750.50)
	remain2.SetLastUpdateAmount(100.00)
	remain2.SetLastUpdateId(2)
	remain2.SetLastUpdateGroup("Withdrawal")
	remain2.SetUpdBy("admin")
	remain2.SetDateActualFrom(time.Now())
	remain2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	Remains = []models.Remain{remain1, remain2}
}

func income() {
	income1 := &models.Income{}
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

	income2 := &models.Income{}
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

	Incomes = []*models.Income{income1, income2}
}

func incomeExpected() {
	incomeExpected1 := &models.IncomeExpected{}
	incomeExpected1.SetIdIncomeEx(1)
	incomeExpected1.SetAmount(150.0)
	incomeExpected1.SetTypeIncome("Salary")
	incomeExpected1.SetIdAccaunt(1)
	incomeExpected1.SetIncomeMonthDate(15)
	incomeExpected1.SetUpdBy("admin")
	incomeExpected1.SetDateActualFrom(time.Now())
	incomeExpected1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	incomeExpected2 := &models.IncomeExpected{}
	incomeExpected2.SetIdIncomeEx(2)
	incomeExpected2.SetAmount(75.0)
	incomeExpected2.SetTypeIncome("Bonus")
	incomeExpected2.SetIdAccaunt(2)
	incomeExpected2.SetIncomeMonthDate(20)
	incomeExpected2.SetUpdBy("admin")
	incomeExpected2.SetDateActualFrom(time.Now())
	incomeExpected2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	IncomesExpected = []*models.IncomeExpected{incomeExpected1, incomeExpected2}
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

func goal() {
	goal1 := models.Goal{}
	goal1.SetIdGoal(1)
	goal1.SetIdAccaunt(1)
	goal1.SetAmount(1000.0)
	goal1.SetDate(time.Now().AddDate(0, 1, 0)) // Дата через 1 месяц
	goal1.SetUpdBy("admin")
	goal1.SetDateActualFrom(time.Now())
	goal1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	goal2 := models.Goal{}
	goal2.SetIdGoal(2)
	goal2.SetIdAccaunt(2)
	goal2.SetAmount(500.0)
	goal2.SetDate(time.Now().AddDate(0, 2, 0)) // Дата через 2 месяца
	goal2.SetUpdBy("admin")
	goal2.SetDateActualFrom(time.Now())
	goal2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	Goals = []models.Goal{goal1, goal2}
}

func cashback() {
	cashback1 := models.Cashback{}
	cashback1.SetIdCashback(1)
	cashback1.SetIdAccaunt(1)
	cashback1.SetBankName("Bank A")
	cashback1.SetCategory("Groceries")
	cashback1.SetPercent(5)
	cashback1.SetUpdBy("admin")
	cashback1.SetDateActualFrom(time.Now())
	cashback1.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	cashback2 := models.Cashback{}
	cashback2.SetIdCashback(2)
	cashback2.SetIdAccaunt(2)
	cashback2.SetBankName("Bank B")
	cashback2.SetCategory("Dining")
	cashback2.SetPercent(10)
	cashback2.SetUpdBy("admin")
	cashback2.SetDateActualFrom(time.Now())
	cashback2.SetDateActualTo(time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC))

	Cashbacks = []models.Cashback{cashback1, cashback2}
}
