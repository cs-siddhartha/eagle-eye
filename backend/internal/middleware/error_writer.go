package middleware

import "net/http"

type ErrorWriter func(
	writer http.ResponseWriter,
	request *http.Request,
	status int,
	code string,
	message string,
)
