package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/your-org/observability-tool/backend/internal/middleware"
)

// Meta carries request-level context that helps clients and operators
// correlate API responses with server logs.
type Meta struct {
	RequestID string `json:"requestId"`
}

// DataResponse provides a stable success envelope for REST resources.
type DataResponse struct {
	Data any  `json:"data"`
	Meta Meta `json:"meta"`
}

// ErrorBody exposes a stable machine-readable code and a safe client message.
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse provides one consistent envelope for all API failures.
type ErrorResponse struct {
	Error ErrorBody `json:"error"`
	Meta  Meta      `json:"meta"`
}

// WriteData serializes a successful REST response with correlation metadata.
func WriteData(logger *slog.Logger, writer http.ResponseWriter, request *http.Request, status int, data any) {
	writeJSON(logger, writer, status, DataResponse{
		Data: data,
		Meta: Meta{RequestID: middleware.RequestIDFromContext(request.Context())},
	})
}

// WriteError serializes a safe API error without leaking internal details.
func WriteError(logger *slog.Logger, writer http.ResponseWriter, request *http.Request, status int, code, message string) {
	writeJSON(logger, writer, status, ErrorResponse{
		Error: ErrorBody{Code: code, Message: message},
		Meta:  Meta{RequestID: middleware.RequestIDFromContext(request.Context())},
	})
}

// writeJSON applies the JSON content type and logs the rare serialization
// failure because a partially written response cannot be recovered.
func writeJSON(logger *slog.Logger, writer http.ResponseWriter, status int, payload any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	if err := json.NewEncoder(writer).Encode(payload); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}
