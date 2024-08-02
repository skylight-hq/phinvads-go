package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// GetValueSetConceptByCodeSystemOID retrieves a row from 'public.value_set_concept' as a [ValueSetConcept].
func GetValueSetConceptsByCodeSystemOID(ctx context.Context, db xo.DB, csOid string) ([]*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByCodesystemoid(ctx, db, csOid)
}

// GetValueSetConceptByID retrieves a row from 'public.value_set_concept' as a [ValueSetConcept].
func GetValueSetConceptByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByID(ctx, db, id)
}

// GetValueSetConceptByValuesetversionid retrieves a row from 'public.value_set_concept' as a [ValueSetConcept].
func GetValueSetConceptByValueSetVersionID(ctx context.Context, db xo.DB, vsvId string) ([]*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByValuesetversionid(ctx, db, vsvId)
}

// GetCodeSystemByValueSetConceptCsOid returns the CodeSystem associated with the [ValueSetConcept]'s (Codesystemoid).
func GetCodeSystemByValueSetConceptCsOid(ctx context.Context, db xo.DB, vsc *xo.ValueSetConcept) (*xo.CodeSystem, error) {
	return vsc.CodeSystem(ctx, db)
}

// ValueSetVersion returns the ValueSetVersion associated with the [ValueSetConcept]'s (Valuesetversionid).
func GetValueSetVersionByVscVsvId(ctx context.Context, db xo.DB, vsc *xo.ValueSetConcept) (*xo.ValueSetVersion, error) {
	return vsc.ValueSetVersion(ctx, db)
}
