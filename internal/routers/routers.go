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
		handlers.IncomeGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdIncome(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeDelete(w, r, loggerConsole, loggerFile, config)
	})

	//income_expected
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedGetByIncomeExpectedId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetByAccountId(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/income_expected/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedDelete(w, r, loggerConsole, loggerFile, config)
	})

	// account
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetByIdAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/account/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountDelete(w, r, loggerConsole, loggerFile, config)
	})

	//expence
	http.HandleFunc("/expence/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByIdExpence(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/group/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByIdGroup(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/title/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByTitle(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByDateBetween(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByAmountBetween(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByAmountLess(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencesGetByAmountMore(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/every/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByRepeat(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencePost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencePut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/expence/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceDelete(w, r, loggerConsole, loggerFile, config)
	})

	//remain
	http.HandleFunc("/remain/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetAll(w, r, loggerConsole, loggerFile, config)
	})

	//goal
	http.HandleFunc("/goal/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByIdGoal(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByIdAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByDateBetween(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountBetween(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountLess(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountMore(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetCurrent(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalPost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalPut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/goal/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalDelete(w, r, loggerConsole, loggerFile, config)
	})

	//cashback
	http.HandleFunc("/cashback/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetAll(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdCashback(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdAccount(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/bank/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByBankName(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/category/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByCategory(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetCurrent(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPost(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPut(w, r, loggerConsole, loggerFile, config)
	})
	http.HandleFunc("/cashback/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackDelete(w, r, loggerConsole, loggerFile, config)
	})
}
