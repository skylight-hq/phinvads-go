package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// GetViewVersionByID retrieves a row from 'public.view_version' as a [ViewVersion].
func GetViewVersionByID(ctx context.Context, db xo.DB, id string) (*xo.ViewVersion, error) {
	return xo.ViewVersionByID(ctx, db, id)
}

// ViewVersionByViewId retrieves a row from 'public.view_version' as a [ViewVersion].
func GetViewVersionByViewId(ctx context.Context, db xo.DB, viewId string) ([]*xo.ViewVersion, error) {
	return xo.ViewVersionByViewid(ctx, db, viewId)
}

// GetViewByViewVersionViewId returns the View associated with the [ViewVersion]'s (Viewid).
func GetViewByViewVersionViewId(ctx context.Context, db xo.DB, vv *xo.ViewVersion) (*xo.View, error) {
	return vv.View(ctx, db)
}
