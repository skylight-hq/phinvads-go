package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/skylight-hq/phinvads-go/internal/models"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}

func (app *application) getAllCodeSystems(w http.ResponseWriter, r *http.Request) {
	codeSystems, err := models.GetAllCodeSystems(r.Context(), app.db)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystems)
}

func (app *application) getCodeSystemByOID(w http.ResponseWriter, r *http.Request) {
	oid := r.PathValue("oid")

	codeSystem, err := models.CodeSystemByOid(r.Context(), app.db, oid)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System %s not found", oid)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystem)
}

func (app *application) getAllViews(w http.ResponseWriter, r *http.Request) {
	views, err := models.GetAllViews(r.Context(), app.db)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(views)
}

func (app *application) getViewByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	view, err := models.ViewByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View not found", id)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(view)
}

func (app *application) getViewVersionByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	viewVersion, err := models.ViewVersionByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", id)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersion)
}

func (app *application) getViewVersionsByViewID(w http.ResponseWriter, r *http.Request) {
	viewId := r.PathValue("viewId")

	viewVersions, err := models.ViewVersionByViewid(r.Context(), app.db, viewId)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", viewId)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersions)
}

func (app *application) getAllCodeSystemConcepts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	codeSystemConcepts, err := models.GetAllCodeSystemConcepts(ctx, app.db)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcepts)
}

func (app *application) getCodeSystemConceptByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := r.PathValue("id")

	codeSystemConcept, err := models.CodeSystemConceptByID(ctx, app.db, id)
	if err != nil {
		if errors.Is(err, models.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System Concept%s not found", id)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcept)
}
