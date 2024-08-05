package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/skylight-hq/phinvads-go/internal/database/models"
	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
	e "github.com/skylight-hq/phinvads-go/internal/errors"
)

func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}

func (app *Application) getAllCodeSystems(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	codeSystems, err := rp.GetAllCodeSystems(r.Context(), app.db)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystems)
}

func (app *Application) getCodeSystemByOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	oid := r.PathValue("oid")

	codeSystem, err := rp.GetCodeSystemByOID(r.Context(), app.db, oid)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
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

func (app *Application) getAllViews(w http.ResponseWriter, r *http.Request) {
	views, err := models.GetAllViews(r.Context(), app.db)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(views)
}

func (app *Application) getViewByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	view, err := models.GetViewByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View %s not found", id)
			http.Error(w, errorString, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(view)
}

func (app *Application) getViewVersionByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	viewVersion, err := models.GetViewVersionByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
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

func (app *Application) getViewVersionsByViewID(w http.ResponseWriter, r *http.Request) {
	viewId := r.PathValue("viewId")

	viewVersions, err := models.GetViewVersionByViewId(r.Context(), app.db, viewId)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
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

func (app *Application) getAllCodeSystemConcepts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	codeSystemConcepts, err := models.GetAllCodeSystemConcepts(ctx, app.db)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcepts)
}

func (app *Application) getCodeSystemConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	ctx := context.Background()

	id := r.PathValue("id")

	codeSystemConcept, err := rp.GetCodeSystemConceptByID(ctx, app.db, id)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
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

func (app *Application) getAllValueSets(w http.ResponseWriter, r *http.Request) {
	valueSets, err := models.GetAllValueSets(r.Context(), app.db)
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSets)
}

// getValueSetByOid can retrieve a resource via either ID or an OID (see models/valueset.xo.go)
func (app *Application) getValueSetByOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("oid")

	valueSet, err := rp.GetValueSetByOID(r.Context(), app.db, id)

	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set %s not found", id)
			http.Error(w, errorString, http.StatusNotFound)
		} else if errors.Is(err, e.ErrBadRequest) {
			app.badRequest(w, r, err)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSet)
}
