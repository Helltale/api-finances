package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func goal(logger *logger.CombinedLogger, config *config.Config) {
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
}
