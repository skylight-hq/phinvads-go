package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api", app.healthcheck)
	mux.HandleFunc("GET /api/code-systems", app.getAllCodeSystems)
	mux.HandleFunc("GET /api/code-systems/{oid}", app.getCodeSystemByOID)
	mux.HandleFunc("GET /api/code-system-concepts", app.getAllCodeSystemConcepts)
	mux.HandleFunc("GET /api/code-system-concepts/{id}", app.getCodeSystemConceptByID)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
