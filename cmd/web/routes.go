package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /api", app.healthcheck)
	mux.HandleFunc("GET /api/code-systems", app.getAllCodeSystems)
	mux.HandleFunc("GET /api/code-systems/{oid}", app.getCodeSystemByOID)

	mux.HandleFunc("POST /api/count", app.count)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
