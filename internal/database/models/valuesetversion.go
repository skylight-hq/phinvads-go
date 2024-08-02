package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// ValueSetVersionByID retrieves a row from 'public.value_set_version' as a [ValueSetVersion].
func GetValueSetVersionByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetVersion, error) {
	return xo.ValueSetVersionByID(ctx, db, id)
}

// GetValueSetVersionByValueSetOID retrieves a row from 'public.value_set_version' as a [ValueSetVersion].
func GetValueSetVersionByValueSetOID(ctx context.Context, db xo.DB, oid string) ([]*xo.ValueSetVersion, error) {
	return xo.ValueSetVersionByValuesetoid(ctx, db, oid)
}

// GetValueSetByVersionOID returns the ValueSet associated with the [ValueSetVersion]'s (Valuesetoid).
func GetValueSetByVersionOID(ctx context.Context, db xo.DB, vsv *xo.ValueSetVersion) (*xo.ValueSet, error) {
	return vsv.ValueSet(ctx, db)
}
