package repository

import (
	"context"
	"database/sql"

	"github.com/skylight-hq/phinvads-go/internal/database/models"
	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// Repository contains all the CRUD methods for all resource types represented in the models package
type Repository struct {
	database xo.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{database: db}
}

// =============================== //
// ====== CodeSystem methods ===== //
// =============================== //
func (r *Repository) GetAllCodeSystems(ctx context.Context, db xo.DB) (*[]xo.CodeSystem, error) {
	return models.GetAllCodeSystems(ctx, db)
}

func (r *Repository) GetCodeSystemByID(ctx context.Context, db xo.DB, id string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByID(ctx, db, id)
}

func (r *Repository) GetCodeSystemByOID(ctx context.Context, db xo.DB, oid string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByOid(ctx, db, oid)
}

// =============================== //
// == CodeSystemConcept methods == //
// =============================== //
func (r *Repository) GetAllCodeSystemConcepts(ctx context.Context, db xo.DB) (*[]xo.CodeSystemConcept, error) {
	return models.GetAllCodeSystemConcepts(ctx, db)
}

func (r *Repository) GetCodeSystemConceptByID(ctx context.Context, db xo.DB, id string) (*xo.CodeSystemConcept, error) {
	return xo.CodeSystemConceptByID(ctx, db, id)
}

func (r *Repository) GetCodeSystemConceptsByOID(ctx context.Context, db xo.DB, oid string) ([]*xo.CodeSystemConcept, error) {
	return xo.CodeSystemConceptByCodesystemoid(ctx, db, oid)
}

func (r *Repository) GetCodeSystemByValueSetConceptCsOid(ctx context.Context, db xo.DB, vsc *xo.ValueSetConcept) (*xo.CodeSystem, error) {
	return vsc.CodeSystem(ctx, db)
}

// =============================== //
// ====== ValueSet methods ======= //
// =============================== //
func (r *Repository) GetAllValueSets(ctx context.Context, db xo.DB) (*[]xo.ValueSet, error) {
	return models.GetAllValueSets(ctx, db)
}

func (r *Repository) GetValueSetByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSet, error) {
	return xo.ValueSetByID(ctx, db, id)
}

func (r *Repository) GetValueSetByOID(ctx context.Context, db xo.DB, oid string) (*xo.ValueSet, error) {
	return xo.ValueSetByOid(ctx, db, oid)
}

func (r *Repository) GetValueSetByVersionOID(ctx context.Context, db xo.DB, vsv *xo.ValueSetVersion) (*xo.ValueSet, error) {
	return vsv.ValueSet(ctx, db)
}

// =============================== //
// ========= View methods ======== //
// =============================== //
func (r *Repository) GetAllViews(ctx context.Context, db xo.DB) (*[]xo.View, error) {
	return models.GetAllViews(ctx, db)
}

func (r *Repository) GetViewByID(ctx context.Context, db xo.DB, id string) (*xo.View, error) {
	return xo.ViewByID(ctx, db, id)
}

func (r *Repository) GetViewByViewVersionId(ctx context.Context, db xo.DB, vv *xo.ViewVersion) (*xo.View, error) {
	return models.GetViewByViewVersionViewId(ctx, db, vv)
}

// =============================== //
// ===== ViewValueSet methods ==== //
// =============================== //
func (r *Repository) GetViewValueSetVersionByVvIdVsvId(ctx context.Context, db xo.DB, viewVersionId, valueSetVersionId string) (*xo.ViewValueSetVersion, error) {
	return models.GetViewValueSetVersionByVvIdVsvId(ctx, db, viewVersionId, valueSetVersionId)
}

// =============================== //
// ===== ViewVersion methods ===== //
// =============================== //
func (r *Repository) GetViewVersionByID(ctx context.Context, db xo.DB, id string) (*xo.ViewVersion, error) {
	return models.GetViewVersionByID(ctx, db, id)
}

func (r *Repository) GetViewVersionByViewId(ctx context.Context, db xo.DB, viewId string) ([]*xo.ViewVersion, error) {
	return models.GetViewVersionByViewId(ctx, db, viewId)
}

func (r *Repository) GetViewVersionByVvsvVvId(ctx context.Context, db xo.DB, vvsv *xo.ViewValueSetVersion) (*xo.ViewVersion, error) {
	return models.GetViewVersionByVvsvVvId(ctx, db, vvsv)
}

// =============================== //
// === ValueSetConcept methods === //
// =============================== //
func (r *Repository) GetValueSetConceptsByCodeSystemOID(ctx context.Context, db xo.DB, csOid string) ([]*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByCodesystemoid(ctx, db, csOid)
}

func (r *Repository) GetValueSetConceptByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByID(ctx, db, id)
}

func (r *Repository) GetValueSetConceptByValueSetVersionID(ctx context.Context, db xo.DB, vsvId string) ([]*xo.ValueSetConcept, error) {
	return xo.ValueSetConceptByValuesetversionid(ctx, db, vsvId)
}

// =============================== //
// ==== ValueSetGroup methods ==== //
// =============================== //
func (r *Repository) GetValueSetGroupByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetGroup, error) {
	return xo.ValueSetGroupByID(ctx, db, id)
}

// =============================== //
// === ValueSetVersion methods === //
// =============================== //

func (r *Repository) GetValueSetVersionByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetVersion, error) {
	return xo.ValueSetVersionByID(ctx, db, id)
}

func (r *Repository) GetValueSetVersionByValueSetOID(ctx context.Context, db xo.DB, oid string) ([]*xo.ValueSetVersion, error) {
	return xo.ValueSetVersionByValuesetoid(ctx, db, oid)
}

func (r *Repository) GetValueSetVersionByVscVsvId(ctx context.Context, db xo.DB, vsc *xo.ValueSetConcept) (*xo.ValueSetVersion, error) {
	return vsc.ValueSetVersion(ctx, db)
}

func (r *Repository) GetValueSetVersionByVvsvVsvId(ctx context.Context, db xo.DB, vvsv *xo.ViewValueSetVersion) (*xo.ValueSetVersion, error) {
	return models.GetValueSetVersionByVvsvVsvId(ctx, db, vvsv)
}
