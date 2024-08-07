package errors

import (
	"fmt"
	"net/http"
)

// Request Errors
type RequestError struct {
	Err    error
	Msg    string
	Method string
	Uri    string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("Request Error: %s", e.Msg)
}

var ErrInvalidId = fmt.Errorf("invalid identifier")

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	// var (
	// 	method = r.Method
	// 	uri    = r.URL.RequestURI()
	// )

	// app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// Database Errors
type DatabaseError struct {
	Err    error
	Msg    string
	Method string
	Id     string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Msg: %s. Err: %s", e.Msg, e.Err)
}

func (e *DatabaseError) NoRows(w http.ResponseWriter, r *http.Request, err error) {
	// var (
	// 	method = r.Method
	// 	uri    = r.URL.RequestURI()
	// )

	// app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	http.Error(w, e.Msg, http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	// var (
	// 	method = r.Method
	// 	uri    = r.URL.RequestURI()
	// )

	// app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// func (app *Application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }
