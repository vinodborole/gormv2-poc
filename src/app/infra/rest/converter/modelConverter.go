package converter

import (
	"app/infra/database/model"
	Restmodel "app/infra/rest/generated/server/go"
)

//GetFitnessAppSwaggerResponse get fitness app swagger response
func GetAppSwaggerResponse(dbFitnessApp *model.Apps) (Restmodel.App, error) {
	var App Restmodel.App
	App.Description = dbFitnessApp.Description.String
	App.Name = dbFitnessApp.Name.String
	App.Port = dbFitnessApp.Port.String
	App.Url = dbFitnessApp.URL.String
	return App, nil
}
