package models

import (
	"context"

	"github.com/skylight-hq/phinvads-go/internal/database/models/xo"
)

func GetCodeSystemByID(ctx context.Context, db xo.DB, id string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByOid(ctx, db, id)
}

func GetCodeSystemByOID(ctx context.Context, db xo.DB, oid string) (*xo.CodeSystem, error) {
	return xo.CodeSystemByOid(ctx, db, oid)
}

func GetAllCodeSystems(ctx context.Context, db xo.DB) (*[]xo.CodeSystem, error) {
	const sqlstr = `SELECT * FROM public.code_system`

	codeSystems := []xo.CodeSystem{}
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cs := xo.CodeSystem{}
		err := rows.Scan(&cs.Oid, &cs.ID, &cs.Name, &cs.Definitiontext, &cs.Status, &cs.Version, &cs.Versiondescription, &cs.Assigningauthorityversionname, &cs.Distributionsourceversionname, &cs.Distributionsourceid, &cs.Assigningauthorityid, &cs.Codesystemcode, &cs.Sourceurl, &cs.Hl70396identifier, &cs.Legacyflag, &cs.Statusdate, &cs.Acquireddate, &cs.Effectivedate, &cs.Expirydate, &cs.Assigningauthorityreleasedate, &cs.Distributionsourcereleasedate, &cs.Sdocreatedate, &cs.Lastrevisiondate)
		if err != nil {
			return nil, err
		}
		codeSystems = append(codeSystems, cs)
	}
	return &codeSystems, nil
}
