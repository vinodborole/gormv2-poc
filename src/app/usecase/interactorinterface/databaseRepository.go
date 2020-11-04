package interactorinterface

import (
	"app/infra/database"
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
	GetApp(appName string) (database.App, error)
	CreateApp(app *database.App, ctx context.Context) error
}
