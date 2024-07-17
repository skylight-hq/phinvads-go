package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// ValueSet represents a row from 'public.value_set'.
type ValueSet struct {
	Oid                      string         `json:"oid"`                      // oid
	ID                       sql.NullString `json:"id"`                       // id
	Name                     sql.NullString `json:"name"`                     // name
	Code                     sql.NullString `json:"code"`                     // code
	Status                   sql.NullString `json:"status"`                   // status
	Definitiontext           sql.NullString `json:"definitiontext"`           // definitiontext
	Scopenotetext            sql.NullString `json:"scopenotetext"`            // scopenotetext
	Assigningauthorityid     sql.NullString `json:"assigningauthorityid"`     // assigningauthorityid
	Legacyflag               sql.NullBool   `json:"legacyflag"`               // legacyflag
	Statusdate               sql.NullTime   `json:"statusdate"`               // statusdate
	Valuesetcreateddate      sql.NullTime   `json:"valuesetcreateddate"`      // valuesetcreateddate
	Valuesetlastrevisiondate sql.NullTime   `json:"valuesetlastrevisiondate"` // valuesetlastrevisiondate
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [ValueSet] exists in the database.
func (vs *ValueSet) Exists() bool {
	return vs._exists
}

// Deleted returns true when the [ValueSet] has been marked for deletion
// from the database.
func (vs *ValueSet) Deleted() bool {
	return vs._deleted
}

// Insert inserts the [ValueSet] to the database.
func (vs *ValueSet) Insert(ctx context.Context, db DB) error {
	switch {
	case vs._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case vs._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.value_set (` +
		`oid, id, name, code, status, definitiontext, scopenotetext, assigningauthorityid, legacyflag, statusdate, valuesetcreateddate, valuesetlastrevisiondate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)`
	// run
	logf(sqlstr, vs.Oid, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate)
	if _, err := db.ExecContext(ctx, sqlstr, vs.Oid, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate); err != nil {
		return logerror(err)
	}
	// set exists
	vs._exists = true
	return nil
}

// Update updates a [ValueSet] in the database.
func (vs *ValueSet) Update(ctx context.Context, db DB) error {
	switch {
	case !vs._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case vs._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.value_set SET ` +
		`id = $1, name = $2, code = $3, status = $4, definitiontext = $5, scopenotetext = $6, assigningauthorityid = $7, legacyflag = $8, statusdate = $9, valuesetcreateddate = $10, valuesetlastrevisiondate = $11 ` +
		`WHERE oid = $12`
	// run
	logf(sqlstr, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate, vs.Oid)
	if _, err := db.ExecContext(ctx, sqlstr, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate, vs.Oid); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [ValueSet] to the database.
func (vs *ValueSet) Save(ctx context.Context, db DB) error {
	if vs.Exists() {
		return vs.Update(ctx, db)
	}
	return vs.Insert(ctx, db)
}

// Upsert performs an upsert for [ValueSet].
func (vs *ValueSet) Upsert(ctx context.Context, db DB) error {
	switch {
	case vs._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.value_set (` +
		`oid, id, name, code, status, definitiontext, scopenotetext, assigningauthorityid, legacyflag, statusdate, valuesetcreateddate, valuesetlastrevisiondate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)` +
		` ON CONFLICT (oid) DO ` +
		`UPDATE SET ` +
		`id = EXCLUDED.id, name = EXCLUDED.name, code = EXCLUDED.code, status = EXCLUDED.status, definitiontext = EXCLUDED.definitiontext, scopenotetext = EXCLUDED.scopenotetext, assigningauthorityid = EXCLUDED.assigningauthorityid, legacyflag = EXCLUDED.legacyflag, statusdate = EXCLUDED.statusdate, valuesetcreateddate = EXCLUDED.valuesetcreateddate, valuesetlastrevisiondate = EXCLUDED.valuesetlastrevisiondate `
	// run
	logf(sqlstr, vs.Oid, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate)
	if _, err := db.ExecContext(ctx, sqlstr, vs.Oid, vs.ID, vs.Name, vs.Code, vs.Status, vs.Definitiontext, vs.Scopenotetext, vs.Assigningauthorityid, vs.Legacyflag, vs.Statusdate, vs.Valuesetcreateddate, vs.Valuesetlastrevisiondate); err != nil {
		return logerror(err)
	}
	// set exists
	vs._exists = true
	return nil
}

// Delete deletes the [ValueSet] from the database.
func (vs *ValueSet) Delete(ctx context.Context, db DB) error {
	switch {
	case !vs._exists: // doesn't exist
		return nil
	case vs._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.value_set ` +
		`WHERE oid = $1`
	// run
	logf(sqlstr, vs.Oid)
	if _, err := db.ExecContext(ctx, sqlstr, vs.Oid); err != nil {
		return logerror(err)
	}
	// set deleted
	vs._deleted = true
	return nil
}

// ValueSetByID retrieves a row from 'public.value_set' as a [ValueSet].
//
// Generated from index 'value_set_id_key'.
func ValueSetByID(ctx context.Context, db DB, id sql.NullString) (*ValueSet, error) {
	// query
	const sqlstr = `SELECT ` +
		`oid, id, name, code, status, definitiontext, scopenotetext, assigningauthorityid, legacyflag, statusdate, valuesetcreateddate, valuesetlastrevisiondate ` +
		`FROM public.value_set ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	vs := ValueSet{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&vs.Oid, &vs.ID, &vs.Name, &vs.Code, &vs.Status, &vs.Definitiontext, &vs.Scopenotetext, &vs.Assigningauthorityid, &vs.Legacyflag, &vs.Statusdate, &vs.Valuesetcreateddate, &vs.Valuesetlastrevisiondate); err != nil {
		return nil, logerror(err)
	}
	return &vs, nil
}

// ValueSetByOid retrieves a row from 'public.value_set' as a [ValueSet].
//
// Generated from index 'value_set_pkey'.
func ValueSetByOid(ctx context.Context, db DB, oid string) (*ValueSet, error) {
	// query
	const sqlstr = `SELECT ` +
		`oid, id, name, code, status, definitiontext, scopenotetext, assigningauthorityid, legacyflag, statusdate, valuesetcreateddate, valuesetlastrevisiondate ` +
		`FROM public.value_set ` +
		`WHERE oid = $1`
	// run
	logf(sqlstr, oid)
	vs := ValueSet{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, oid).Scan(&vs.Oid, &vs.ID, &vs.Name, &vs.Code, &vs.Status, &vs.Definitiontext, &vs.Scopenotetext, &vs.Assigningauthorityid, &vs.Legacyflag, &vs.Statusdate, &vs.Valuesetcreateddate, &vs.Valuesetlastrevisiondate); err != nil {
		return nil, logerror(err)
	}
	return &vs, nil
}
