package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

// ValueSetGroup retrieves a row from 'public.value_set_group' as a [ValueSetGroup].
func GetValueSetGroupByID(ctx context.Context, db xo.DB, id string) (*xo.ValueSetGroup, error) {
	return xo.ValueSetGroupByID(ctx, db, id)
}
