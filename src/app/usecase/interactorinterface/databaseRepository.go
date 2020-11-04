package interactorinterface

import (
	"app/infra/database/model"
	"context"
	"gorm.io/gorm"
)

//DatabaseRepository database repository
type DatabaseRepository interface {
	//Transaction Operations
	OpenTransaction() (*gorm.DB, error)
	CommitTransaction(transaction *gorm.DB) error
	RollBackTransaction(transaction *gorm.DB) error

	//App Info
	GetApp(appName string) (model.Apps, error)
	CreateApp(app *model.Apps, ctx context.Context) error

	//Fabric
	BulkInsertFabrics() error
}
