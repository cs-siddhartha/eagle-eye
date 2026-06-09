package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const requestIDHeader = "X-Request-ID"

type requestIDContextKey struct{}

// RequestID assigns or preserves a request identifier so clients can
// correlate each response with future structured server logs.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestID := request.Header.Get(requestIDHeader)
		if requestID == "" {
			requestID = newRequestID()
		}

		writer.Header().Set(requestIDHeader, requestID)
		contextWithRequestID := context.WithValue(
			request.Context(),
			requestIDContextKey{},
			requestID,
		)

		next.ServeHTTP(writer, request.WithContext(contextWithRequestID))
	})
}

func RequestIDFromContext(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDContextKey{}).(string)
	return requestID
}

func newRequestID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "request-id-unavailable"
	}

	return hex.EncodeToString(bytes)
}
