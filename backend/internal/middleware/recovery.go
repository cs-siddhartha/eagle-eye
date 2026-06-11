package middleware

import (
	"log/slog"
	"net/http"
)

func Recovery(logger *slog.Logger, writeError ErrorWriter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				logger.Error(
					"panic recovered while handling request",
					"requestId", RequestIDFromContext(request.Context()),
					"method", request.Method,
					"path", request.URL.Path,
					"panic", recovered,
				)
				writeError(
					writer,
					request,
					http.StatusInternalServerError,
					"INTERNAL_ERROR",
					"An unexpected error occurred",
				)
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
