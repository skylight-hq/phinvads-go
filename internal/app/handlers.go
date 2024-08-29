package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/skylight-hq/phinvads-go/internal/database/models"
	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
	customErrors "github.com/skylight-hq/phinvads-go/internal/errors"
	"github.com/skylight-hq/phinvads-go/internal/ui/components"
)

func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: OK"))
}

func (app *Application) getAllCodeSystems(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	codeSystems, err := rp.GetAllCodeSystems(r.Context())
	if err != nil {
		if errors.Is(err, xo.ErrDoesNotExist) {
			http.NotFound(w, r)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
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
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var codeSystem *xo.CodeSystem
	if id_type == "oid" {
		codeSystem, err = rp.GetCodeSystemByOID(r.Context(), id)
	} else {
		codeSystem, err = rp.GetCodeSystemByID(r.Context(), id)
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
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystem)
}

func (app *Application) getAllViews(w http.ResponseWriter, r *http.Request) {
	views, err := models.GetAllViews(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(views)
}

func (app *Application) getViewByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	view, err := rp.GetViewByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(view)
}

func (app *Application) getViewVersionByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	viewVersion, err := rp.GetViewVersionByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersion)
}

func (app *Application) getViewVersionsByViewID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	viewId := r.PathValue("viewId")

	viewVersions, err := rp.GetViewVersionByViewId(r.Context(), viewId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: View Version %s not found", viewId)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getViewVersionsByViewID",
				Id:     viewId,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(viewVersions)
}

func (app *Application) getAllCodeSystemConcepts(w http.ResponseWriter, r *http.Request) {
	codeSystemConcepts, err := models.GetAllCodeSystemConcepts(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcepts)
}

func (app *Application) getCodeSystemConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")

	codeSystemConcept, err := rp.GetCodeSystemConceptByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Code System Concept%s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getCodeSystemConceptByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(codeSystemConcept)
}

func (app *Application) getAllValueSets(w http.ResponseWriter, r *http.Request) {
	valueSets, err := models.GetAllValueSets(r.Context(), app.db)
	if err != nil {
		customErrors.ServerError(w, r, err, app.logger)
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
		customErrors.BadRequest(w, r, err, app.logger)
		return
	}

	var valueSet *xo.ValueSet
	if id_type == "oid" {
		valueSet, err = rp.GetValueSetByOID(r.Context(), id)
	} else {
		valueSet, err = rp.GetValueSetByID(r.Context(), id)
	}

	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSet)
}

func (app *Application) getValueSetVersionsByValueSetOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	oid := r.PathValue("oid")

	valueSetVersions, err := rp.GetValueSetVersionByValueSetOID(r.Context(), oid)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Versions with Value Set OID %s not found", oid)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetVersionsByValueSetOID",
				Id:     oid,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSetVersions)
}

func (app *Application) getValueSetVersionByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	id := r.PathValue("id")

	valueSetVersion, err := rp.GetValueSetVersionByID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Version %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetVersionsByValueSetOID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSetVersion)
}

func (app *Application) getValueSetConceptByID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("id")

	valueSetConcept, err := rp.GetValueSetConceptByID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: Value Set Concept %s not found", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptByID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}
		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSetConcept)
}

func (app *Application) getValueSetConceptsByVersionID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	id := r.PathValue("valueSetVersionId")

	valueSetConcepts, err := rp.GetValueSetConceptByValueSetVersionID(r.Context(), id)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: No Value Set Concepts found for version %s", id)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptsByVersionID",
				Id:     id,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSetConcepts)
}

func (app *Application) getValueSetConceptsByCodeSystemOID(w http.ResponseWriter, r *http.Request) {
	rp := app.repository
	oid := r.PathValue("codeSystemOid")

	valueSetConcepts, err := rp.GetValueSetConceptsByCodeSystemOID(r.Context(), oid)
	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := fmt.Sprintf("Error: No Value Set Concepts found for CodeSystem %s", oid)
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "getValueSetConceptsByCodeSystemOID",
				Id:     oid,
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(valueSetConcepts)
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	component.Render(r.Context(), w)
}

func (app *Application) handleBannerToggle(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")
	var path string
	if action == "close" {
		path = "internal/ui/components/banner-closed.html"

	} else {
		path = "internal/ui/components/banner-open.html"
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
