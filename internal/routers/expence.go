package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func expence(logger *logger.CombinedLogger, config config.Config) {
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
}
