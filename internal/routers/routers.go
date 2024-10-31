package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func Init(logger *logger.CombinedLogger, config config.Config) {
	//income
	http.HandleFunc("/income/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetAll(w, r, logger, config)
	})
	http.HandleFunc("/income/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdIncome(w, r, logger, config)
	})
	http.HandleFunc("/income/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/income/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePost(w, r, logger, config)
	})
	http.HandleFunc("/income/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomePut(w, r, logger, config)
	})
	http.HandleFunc("/income/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeDelete(w, r, logger, config)
	})

	//income_expected
	http.HandleFunc("/income_expected/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetAll(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedGetByIncomeExpectedId(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomesExpectedGetByAccountId(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPost(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedPut(w, r, logger, config)
	})
	http.HandleFunc("/income_expected/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IncomeExpectedDelete(w, r, logger, config)
	})

	// account
	http.HandleFunc("/account/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetAll(w, r, logger, config)
	})
	http.HandleFunc("/account/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/account/new/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPost(w, r, logger, config)
	})
	http.HandleFunc("/account/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountPut(w, r, logger, config)
	})
	http.HandleFunc("/account/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AccountDelete(w, r, logger, config)
	})

	//expence
	http.HandleFunc("/expence/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetAll(w, r, logger, config)
	})
	http.HandleFunc("/expence/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByIdExpence(w, r, logger, config)
	})
	http.HandleFunc("/expence/group/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByIdGroup(w, r, logger, config)
	})
	http.HandleFunc("/expence/title/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByTitle(w, r, logger, config)
	})
	http.HandleFunc("/expence/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByDateBetween(w, r, logger, config)
	})
	http.HandleFunc("/expence/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByAmountBetween(w, r, logger, config)
	})
	http.HandleFunc("/expence/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByAmountLess(w, r, logger, config)
	})
	http.HandleFunc("/expence/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencesGetByAmountMore(w, r, logger, config)
	})
	http.HandleFunc("/expence/every/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceGetByRepeat(w, r, logger, config)
	})
	http.HandleFunc("/expence/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencePost(w, r, logger, config)
	})
	http.HandleFunc("/expence/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpencePut(w, r, logger, config)
	})
	http.HandleFunc("/expence/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ExpenceDelete(w, r, logger, config)
	})

	//remain
	http.HandleFunc("/remain/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetAll(w, r, logger, config)
	})
	http.HandleFunc("/remain/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetByIdRemain(w, r, logger, config)
	})
	http.HandleFunc("/remain/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/remain/last/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetByIdLastEntry(w, r, logger, config)
	})
	http.HandleFunc("/remain/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainGetByDateBetween(w, r, logger, config)
	})
	http.HandleFunc("/remain/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainPost(w, r, logger, config)
	})
	http.HandleFunc("/remain/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainPut(w, r, logger, config)
	})
	http.HandleFunc("/remain/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemainDelete(w, r, logger, config)
	})

	//goal
	http.HandleFunc("/goal/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetAll(w, r, logger, config)
	})
	http.HandleFunc("/goal/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByIdGoal(w, r, logger, config)
	})
	http.HandleFunc("/goal/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/goal/date/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByDateBetween(w, r, logger, config)
	})
	http.HandleFunc("/goal/amount/between/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountBetween(w, r, logger, config)
	})
	http.HandleFunc("/goal/amount/less/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountLess(w, r, logger, config)
	})
	http.HandleFunc("/goal/amount/more/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetByAmountMore(w, r, logger, config)
	})
	http.HandleFunc("/goal/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalGetCurrent(w, r, logger, config)
	})
	http.HandleFunc("/goal/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalPost(w, r, logger, config)
	})
	http.HandleFunc("/goal/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalPut(w, r, logger, config)
	})
	http.HandleFunc("/goal/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GoalDelete(w, r, logger, config)
	})

	//cashback
	http.HandleFunc("/cashback/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetAll(w, r, logger, config)
	})
	http.HandleFunc("/cashback/id/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdCashback(w, r, logger, config)
	})
	http.HandleFunc("/cashback/account/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByIdAccount(w, r, logger, config)
	})
	http.HandleFunc("/cashback/bank/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByBankName(w, r, logger, config)
	})
	http.HandleFunc("/cashback/category/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetByCategory(w, r, logger, config)
	})
	http.HandleFunc("/cashback/current", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackGetCurrent(w, r, logger, config)
	})
	http.HandleFunc("/cashback/new", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPost(w, r, logger, config)
	})
	http.HandleFunc("/cashback/update/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackPut(w, r, logger, config)
	})
	http.HandleFunc("/cashback/delete/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CashbackDelete(w, r, logger, config)
	})
}
