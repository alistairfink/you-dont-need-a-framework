package handlers

import (
	"context"
	"log/slog"
	"net/http"
)

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Handler(
	routeInfo route,
	handler func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.Default()
		headers := r.Header.Clone()
		headers.Del("Authorization")
		logger = logger.With(
			slog.Any("context",
				slog.GroupValue(
					slog.String("Method", r.Method),
					slog.String("URL", r.URL.Path),
					slog.Any("Headers", headers),
					slog.String("User-Agent", r.UserAgent()),
				),
			),
		)

		ctx := context.WithValue(r.Context(), "logger", logger)
		handler(w, r.WithContext(ctx))
	}
}
