package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// ValueSetGroup represents a row from 'public.value_set_group'.
type ValueSetGroup struct {
	ID              string         `json:"id"`              // id
	Name            sql.NullString `json:"name"`            // name
	Descriptiontext sql.NullString `json:"descriptiontext"` // descriptiontext
	Status          sql.NullString `json:"status"`          // status
	Groupnotes      sql.NullString `json:"groupnotes"`      // groupnotes
	Statusdate      sql.NullTime   `json:"statusdate"`      // statusdate
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [ValueSetGroup] exists in the database.
func (vsg *ValueSetGroup) Exists() bool {
	return vsg._exists
}

// Deleted returns true when the [ValueSetGroup] has been marked for deletion
// from the database.
func (vsg *ValueSetGroup) Deleted() bool {
	return vsg._deleted
}

// Insert inserts the [ValueSetGroup] to the database.
func (vsg *ValueSetGroup) Insert(ctx context.Context, db DB) error {
	switch {
	case vsg._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case vsg._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.value_set_group (` +
		`id, name, descriptiontext, status, groupnotes, statusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`
	// run
	logf(sqlstr, vsg.ID, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate)
	if _, err := db.ExecContext(ctx, sqlstr, vsg.ID, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate); err != nil {
		return logerror(err)
	}
	// set exists
	vsg._exists = true
	return nil
}

// Update updates a [ValueSetGroup] in the database.
func (vsg *ValueSetGroup) Update(ctx context.Context, db DB) error {
	switch {
	case !vsg._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case vsg._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.value_set_group SET ` +
		`name = $1, descriptiontext = $2, status = $3, groupnotes = $4, statusdate = $5 ` +
		`WHERE id = $6`
	// run
	logf(sqlstr, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate, vsg.ID)
	if _, err := db.ExecContext(ctx, sqlstr, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate, vsg.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [ValueSetGroup] to the database.
func (vsg *ValueSetGroup) Save(ctx context.Context, db DB) error {
	if vsg.Exists() {
		return vsg.Update(ctx, db)
	}
	return vsg.Insert(ctx, db)
}

// Upsert performs an upsert for [ValueSetGroup].
func (vsg *ValueSetGroup) Upsert(ctx context.Context, db DB) error {
	switch {
	case vsg._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.value_set_group (` +
		`id, name, descriptiontext, status, groupnotes, statusdate` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, descriptiontext = EXCLUDED.descriptiontext, status = EXCLUDED.status, groupnotes = EXCLUDED.groupnotes, statusdate = EXCLUDED.statusdate `
	// run
	logf(sqlstr, vsg.ID, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate)
	if _, err := db.ExecContext(ctx, sqlstr, vsg.ID, vsg.Name, vsg.Descriptiontext, vsg.Status, vsg.Groupnotes, vsg.Statusdate); err != nil {
		return logerror(err)
	}
	// set exists
	vsg._exists = true
	return nil
}

// Delete deletes the [ValueSetGroup] from the database.
func (vsg *ValueSetGroup) Delete(ctx context.Context, db DB) error {
	switch {
	case !vsg._exists: // doesn't exist
		return nil
	case vsg._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.value_set_group ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, vsg.ID)
	if _, err := db.ExecContext(ctx, sqlstr, vsg.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	vsg._deleted = true
	return nil
}

// ValueSetGroupByID retrieves a row from 'public.value_set_group' as a [ValueSetGroup].
//
// Generated from index 'value_set_group_pkey'.
func ValueSetGroupByID(ctx context.Context, db DB, id string) (*ValueSetGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, descriptiontext, status, groupnotes, statusdate ` +
		`FROM public.value_set_group ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	vsg := ValueSetGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&vsg.ID, &vsg.Name, &vsg.Descriptiontext, &vsg.Status, &vsg.Groupnotes, &vsg.Statusdate); err != nil {
		return nil, logerror(err)
	}
	return &vsg, nil
}
