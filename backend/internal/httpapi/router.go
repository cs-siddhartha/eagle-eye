package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/your-org/observability-tool/backend/internal/middleware"
)

func NewHandler(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /readyz", readinessHandler(logger))
	mux.HandleFunc("/", notFoundHandler(logger))

	return middleware.RequestID(mux)
}

func readinessHandler(logger *slog.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		WriteData(logger, writer, request, http.StatusOK, map[string]string{
			"status": "ready",
		})
	}
}

func notFoundHandler(logger *slog.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		WriteError(
			logger,
			writer,
			request,
			http.StatusNotFound,
			"NOT_FOUND",
			"Resource not found",
		)
	}
}
