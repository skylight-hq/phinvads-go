package app

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api", app.healthcheck)

	mux.HandleFunc("GET /api/code-systems", app.getAllCodeSystems)
	mux.HandleFunc("GET /api/code-systems/{id}", app.getCodeSystemByID)

	mux.HandleFunc("GET /api/code-system-concepts", app.getAllCodeSystemConcepts)
	mux.HandleFunc("GET /api/code-system-concepts/{id}", app.getCodeSystemConceptByID)

	mux.HandleFunc("GET /api/value-sets", app.getAllValueSets)
	mux.HandleFunc("GET /api/value-sets/{id}", app.getValueSetByID)
	mux.HandleFunc("GET /api/value-sets/{oid}/versions", app.getValueSetVersionsByValueSetOID)

	mux.HandleFunc("GET /api/value-set-versions/{id}", app.getValueSetVersionByID)

	mux.HandleFunc("GET /api/views", app.getAllViews)
	mux.HandleFunc("GET /api/views/{id}", app.getViewByID)

	mux.HandleFunc("GET /api/view-versions/{id}", app.getViewVersionByID)
	mux.HandleFunc("GET /api/view-versions-by-view/{viewId}", app.getViewVersionsByViewID)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
