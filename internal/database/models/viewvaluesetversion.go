package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// GetViewValueSetVersionByVvIdVsvId retrieves a row from 'public.view_value_set_version' as a [ViewValueSetVersion].
func GetViewValueSetVersionByVvIdVsvId(ctx context.Context, db xo.DB, viewVersionId, valueSetVersionId string) (*xo.ViewValueSetVersion, error) {
	return xo.ViewValueSetVersionByViewversionidValuesetversionid(ctx, db, viewVersionId, valueSetVersionId)
}

// GetValueSetVersionByVvsvVsvId returns the ValueSetVersion associated with the [ViewValueSetVersion]'s (Valuesetversionid).
func GetValueSetVersionByVvsvVsvId(ctx context.Context, db xo.DB, vvsv *xo.ViewValueSetVersion) (*xo.ValueSetVersion, error) {
	return vvsv.ValueSetVersion(ctx, db)
}

// GetViewVersionByVvsvVvId returns the ViewVersion associated with the [ViewValueSetVersion]'s (Viewversionid).
func GetViewVersionByVvsvVvId(ctx context.Context, db xo.DB, vvsv *xo.ViewValueSetVersion) (*xo.ViewVersion, error) {
	return vvsv.ViewVersion(ctx, db)
}
