package middleware

import (
	"mime"
	"net/http"
)

// LimitRequestBody rejects known oversized payloads immediately and caps body
// reads so streaming or inaccurate clients cannot consume unbounded memory.
func LimitRequestBody(maxBytes int64, writeError ErrorWriter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.ContentLength > maxBytes {
			writeError(
				writer,
				request,
				http.StatusRequestEntityTooLarge,
				"REQUEST_TOO_LARGE",
				"Request body exceeds the allowed size",
			)
			return
		}

		request.Body = http.MaxBytesReader(writer, request.Body, maxBytes)
		next.ServeHTTP(writer, request)
	})
}

// RequireJSON rejects body-bearing write requests that do not declare JSON,
// preventing handlers from guessing how an incoming payload should be parsed.
func RequireJSON(writeError ErrorWriter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !requestCanCarryJSON(request) {
			next.ServeHTTP(writer, request)
			return
		}

		contentType := request.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil || mediaType != "application/json" {
			writeError(
				writer,
				request,
				http.StatusUnsupportedMediaType,
				"UNSUPPORTED_MEDIA_TYPE",
				"Content-Type must be application/json",
			)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

// requestCanCarryJSON limits content-type enforcement to write methods with a
// body so empty operational and read requests remain valid.
func requestCanCarryJSON(request *http.Request) bool {
	switch request.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return request.Body != nil && request.Body != http.NoBody
	default:
		return false
	}
}
