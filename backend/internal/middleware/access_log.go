package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (recorder *responseRecorder) WriteHeader(status int) {
	if recorder.status != 0 {
		return
	}

	recorder.status = status
	recorder.ResponseWriter.WriteHeader(status)
}

func (recorder *responseRecorder) Write(payload []byte) (int, error) {
	if recorder.status == 0 {
		recorder.WriteHeader(http.StatusOK)
	}

	written, err := recorder.ResponseWriter.Write(payload)
	recorder.bytes += written
	return written, err
}

// Unwrap allows net/http response controllers to reach optional capabilities
// implemented by the original response writer.
func (recorder *responseRecorder) Unwrap() http.ResponseWriter {
	return recorder.ResponseWriter
}

// AccessLog emits one structured record per request for latency, status, and
// request-ID correlation during debugging and operations.
func AccessLog(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startedAt := time.Now()
		recorder := &responseRecorder{ResponseWriter: writer}

		next.ServeHTTP(recorder, request)

		status := recorder.status
		if status == 0 {
			status = http.StatusOK
		}

		logger.Info(
			"http request completed",
			"requestId", RequestIDFromContext(request.Context()),
			"method", request.Method,
			"path", request.URL.Path,
			"status", status,
			"responseBytes", recorder.bytes,
			"durationMs", time.Since(startedAt).Milliseconds(),
			"remoteAddress", request.RemoteAddr,
		)
	})
}
