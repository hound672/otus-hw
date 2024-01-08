package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(rw, r.ProtoMajor)

		requestTime := time.Now()
		next.ServeHTTP(ww, r)
		latency := time.Since(requestTime)

		logger.Info(
			"Request",
			"Client IP", r.RemoteAddr,
			"Request time", requestTime,
			"Method", r.Method,
			"Path", r.URL.Path,
			"HTTP Version", r.Proto,
			"Response code", ww.Status(),
			"Latency", latency,
			"UserAgent", r.UserAgent(),
		)
	})
}
