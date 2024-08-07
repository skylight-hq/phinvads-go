package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/skylight-hq/phinvads-go/internal/database/models"
	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
	customErrors "github.com/skylight-hq/phinvads-go/internal/errors"
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
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystems)
}

func (app *Application) getCodeSystemByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err)
		return
	}

	var codeSystem *xo.CodeSystem
	if id_type == "oid" {
		codeSystem, err = rp.GetCodeSystemByOID(r.Context(), app.db, id)
	} else {
		codeSystem, err = rp.GetCodeSystemByID(r.Context(), app.db, id)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemById",
				Id:     id,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystem)
}

func (app *Application) getAllViews(w http.ResponseWriter, r *http.Request) {
	views, err := models.GetAllViews(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(views)
}

func (app *Application) getViewByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	view, err := rp.GetViewByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(view)
}

func (app *Application) getViewVersionByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	viewVersion, err := rp.GetViewVersionByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersion)
}

func (app *Application) getViewVersionsByViewID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	viewId := r.PathValue("viewId")

	viewVersions, err := rp.GetViewVersionByViewId(r.Context(), app.db, viewId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", viewId)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionsByViewID",
				Id:     viewId,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersions)
}

func (app *Application) getAllCodeSystemConcepts(w http.ResponseWriter, r *http.Request) {
	codeSystemConcepts, err := models.GetAllCodeSystemConcepts(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcepts)
}

func (app *Application) getCodeSystemConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")

	codeSystemConcept, err := rp.GetCodeSystemConceptByID(r.Context(), app.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System Concept%s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemConceptByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcept)
}

func (app *Application) getAllValueSets(w http.ResponseWriter, r *http.Request) {
	valueSets, err := models.GetAllValueSets(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSets)
}

func (app *Application) getValueSetByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")
	id_type, err := determineIdType(id)
	if err != nil {
		customErrors.BadRequest(w, r, err)
		return
	}

	var valueSet *xo.ValueSet
	if id_type == "oid" {
		valueSet, err = rp.GetValueSetByOID(r.Context(), app.db, id)
	} else {
		valueSet, err = rp.GetValueSetByID(r.Context(), app.db, id)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err)
		} else {
			customErrors.ServerError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSet)
}
