package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/skylight-hq/phinvads-go/internal/database/models"
	"github.com/skylight-hq/phinvads-go/internal/database/models/repository"
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

func (app *Application) getAllHotTopics(w http.ResponseWriter, r *http.Request) {
	rp := app.repository

	hotTopics, err := rp.GetAllHotTopics(r.Context())

	if err != nil {
		var (
			method = r.Method
			uri    = r.URL.RequestURI()
		)
		if errors.Is(err, sql.ErrNoRows) {
			errorString := "Error: No Hot Topics found"
			dbErr := &customErrors.DatabaseError{
				Err:    err,
				Msg:    errorString,
				Method: "home: Get Hot Topics",
			}
			dbErr.NoRows(w, r, err, app.logger)
		} else {
			customErrors.ServerError(w, r, err, app.logger)
		}

		app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
		return
	}

	for _, t := range *hotTopics {
		// skip sending system config to the frontend
		if t.HotTopicName == "SYSTEM CONFIG" {
			continue
		}
		// format the sequence id to align with the uswds js controls
		divId := fmt.Sprintf("m-a%s", strconv.Itoa(t.Seq))

		component := components.HotTopic(t.HotTopicName, t.HotTopicDescription, divId, t.HotTopicID.String())
		component.Render(r.Context(), w)
	}
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	component.Render(r.Context(), w)
}

func (app *Application) formSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	searchTerm := r.Form["search"][0]
	searchType := r.Form["options"][0]

	search(w, r, app.repository, searchTerm, searchType, app.logger)
}

func search(w http.ResponseWriter, r *http.Request, rp *repository.Repository, searchTerm, searchType string, logger *slog.Logger) {
	var result = &models.CodeSystemResultRow{}

	// retrieve code system
	codeSystem, err := rp.GetCodeSystemsByLikeOID(r.Context(), searchTerm)
	fmt.Println(codeSystem, err)
	if err != nil || len(*codeSystem) < 1 {
		if err == nil {
			err = sql.ErrNoRows
		}
		handleError(err, searchTerm, logger, w, r)
	}

	for _, cs := range *codeSystem {
		result.CodeSystems = append(result.CodeSystems, &cs)
	}
	result.CodeSystemsCount = strconv.Itoa(len(result.CodeSystems))

	// // retrieve concepts that are part of that code system
	// concepts, err := rp.GetCodeSystemConceptsByCodeSystemOID(r.Context(), app.db, codeSystem)
	// for _, csc := range *concepts {
	// 	result.CodeSystems = append(result.CodeSystems, &csc)
	// }
	// result.CodeSystemConcepts = concepts

	// for now
	result.CodeSystemConceptsCount = strconv.Itoa(0)

	// for now
	result.ValueSetsCount = strconv.Itoa(0)

	w.Header().Set("HX-Push-Url", fmt.Sprintf("/search?type=%s&input=%s", searchType, searchTerm))

	component := components.SearchResults(true, "Search", searchTerm, result)
	component.Render(r.Context(), w)
}

func (app *Application) directSearch(w http.ResponseWriter, r *http.Request) {
	searchType := r.URL.Query().Get("type")
	searchTerm := r.URL.Query().Get("input")
	search(w, r, app.repository, searchTerm, searchType, app.logger)
}

func (app *Application) handleBannerToggle(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")
	component := components.UsaBanner(action)
	component.Render(r.Context(), w)
}

func handleError(err error, input string, logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	if errors.Is(err, sql.ErrNoRows) {
		errorString := fmt.Sprintf("Error: Code System %s not found", input)
		dbErr := &customErrors.DatabaseError{
			Err:    err,
			Msg:    errorString,
			Method: "getCodeSystemById",
			Id:     input,
		}
		component := components.Error("Search", dbErr.Msg)
		component.Render(r.Context(), w)
		dbErr.NoRows(w, r, err, logger)
	} else {
		customErrors.ServerError(w, r, err, logger)
		component := components.Error("search", err.Error())
		component.Render(r.Context(), w)
	}
}
