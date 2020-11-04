package gateway

import (
	"app/infra/database/model"
	"database/sql"
)

func (dbRepo *DatabaseRepository) BulkInsertFabrics() error {
	var fabrics = []model.Fabrics{
		{Name: sql.NullString{String: "fabric1", Valid: true}}, {Name: sql.NullString{String: "fabric2", Valid: true}}, {Name: sql.NullString{String: "fabric3", Valid: true}}}
	dbRepo.GetDBHandle().Create(&fabrics)
	return nil
}
