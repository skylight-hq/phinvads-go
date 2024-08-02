package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

func GetValueSetByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSet, error) {
	return xo.ValueSetByID(ctx, db, id)
}

func GetValueSetByOID(ctx context.Context, db xo.DB, oid string) (*xo.ValueSet, error) {
	return xo.ValueSetByOid(ctx, db, oid)
}

// All retrieves all rows from 'public.value_set'
func GetAllValueSets(ctx context.Context, db xo.DB) (*[]xo.ValueSet, error) {
	const sqlstr = `SELECT * FROM public.value_set`
	valueSets := []xo.ValueSet{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		vs := xo.ValueSet{}
		err := rows.Scan(&vs.Oid, &vs.ID, &vs.Name, &vs.Code, &vs.Status, &vs.Definitiontext, &vs.Scopenotetext, &vs.Assigningauthorityid, &vs.Legacyflag, &vs.Statusdate)
		if err != nil {
			return nil, err
		}
		valueSets = append(valueSets, vs)
	}
	return &valueSets, nil
}
