package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/your-org/observability-tool/backend/internal/middleware"
)

func NewHandler(logger *slog.Logger, maxRequestBytes int64) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/readyz", readinessHandler(logger))
	mux.HandleFunc("/", notFoundHandler(logger))

	writeError := func(
		writer http.ResponseWriter,
		request *http.Request,
		status int,
		code string,
		message string,
	) {
		WriteError(logger, writer, request, status, code, message)
	}

	var handler http.Handler = mux
	handler = middleware.RequireJSON(writeError, handler)
	handler = middleware.LimitRequestBody(maxRequestBytes, writeError, handler)
	handler = middleware.Recovery(logger, writeError, handler)
	handler = middleware.AccessLog(logger, handler)
	handler = middleware.RequestID(handler)

	return handler
}

// unsupported methods with the standard REST error envelope.
func readinessHandler(logger *slog.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.Header().Set("Allow", http.MethodGet)
			WriteError(
				logger,
				writer,
				request,
				http.StatusMethodNotAllowed,
				"METHOD_NOT_ALLOWED",
				"Method not allowed",
			)
			return
		}

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
