package handlers

import (
	"log/slog"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		w,
		http.StatusOK,
	}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type ResponseMiddleware struct {
}

func NewResponseHeaderMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

func (m *ResponseMiddleware) Handler(
	routeInfo route,
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)
		handler(lrw, r)
		if routeInfo.logResponseCode {
			logger, _ := r.Context().Value("logger").(*slog.Logger)
			logger.Info("finished request", slog.Any("status code", lrw.statusCode))
		}
	}
}
