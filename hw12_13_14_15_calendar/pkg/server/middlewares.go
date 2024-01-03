package server

import (
	"net/http"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger.Info("Request", "URL", r.URL.Path)
		next.ServeHTTP(rw, r)
	})
}
