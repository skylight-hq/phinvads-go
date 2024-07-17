package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api", app.healthcheck)
	mux.HandleFunc("GET /api/code-systems", app.getAllCodeSystems)

	return mux
}
