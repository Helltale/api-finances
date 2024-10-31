package routers

import (
	"net/http"

	"github.com/helltale/api-finances/config"
	"github.com/helltale/api-finances/internal/handlers"
	"github.com/helltale/api-finances/internal/logger"
)

func remain(logger *logger.CombinedLogger, config config.Config) {
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
}
