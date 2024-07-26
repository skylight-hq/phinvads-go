package main

import (
	"context"
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
	ctx := context.Background()

	codeSystems, err := models.GetAllCodeSystems(ctx, app.db)
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
	ctx := context.Background()

	oid := r.PathValue("oid")

	codeSystems, err := models.CodeSystemByOid(ctx, app.db, oid)
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

	json.NewEncoder(w).Encode(codeSystems)
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
