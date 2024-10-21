package routers

import (
	"log/slog"
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
)

func Init(loggerConsole *slog.Logger, loggerFile *slog.Logger, config config.Config) {
	//income
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllIncomes(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIncomeById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIncomesByAccountId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostIncome(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutIncome(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteIncome(w, r, loggerConsole, loggerFile, config)
	})

	//income_expected
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIncomesExpectedAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIncomeExpectedById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetIncomesExpectedByAccountId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostIncomeExpected(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutIncomeExpected(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteIncomeExpected(w, r, loggerConsole, loggerFile, config)
	})

	// account
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllAccounts(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAccountById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteAccount(w, r, loggerConsole, loggerFile, config)
	})

	//expence
	http.HandleFunc("/expence/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllExpences(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpenceById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/group/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByGroupId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/title/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByTitle(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByDateRange(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByAmountRange(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByMaxAmount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByMinAmount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/every/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetExpencesByRepeatType(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostExpence(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutExpence(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteExpence(w, r, loggerConsole, loggerFile, config)
	})

	//remain
	http.HandleFunc("/remain/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllRemains(w, r, loggerConsole, loggerFile, config)
	})

	//goal
	http.HandleFunc("/goal/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllGoals(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsByAccountId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsByDateRange(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsByAmountRange(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsByMaxAmount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsByMinAmount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoalsCurrent(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostGoal(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutGoal(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteGoal(w, r, loggerConsole, loggerFile, config)
	})

	//cashback
	http.HandleFunc("/cashback/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllCashbacks(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCashbackById(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCashbacksByAccountId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/bank/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCashbacksByBankName(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/category/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCashbacksByCategory(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCashbacksCurrent(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostCashback(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PutCashback(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteCashback(w, r, loggerConsole, loggerFile, config)
	})
}
