package gateway

import (
	"app/infra/database/model"
	"context"
	"fmt"
)

//GetApp get app info
func (dbRepo *DatabaseRepository) GetApp(appName string) (model.Apps, error) {
	var fitnessApp model.Apps
	err := dbRepo.GetDBHandle().Where("name=?", appName).First(&fitnessApp).Error
	if err != nil {
		return model.Apps{}, err
	}
	return fitnessApp, nil
}

//CreateApp create App
func (dbRepo *DatabaseRepository) CreateApp(app *model.Apps, ctx context.Context) error {
	fmt.Println("creating record ", app)
	err := dbRepo.GetDBHandle().Model(model.Apps{}).Create(&app).Error
	fmt.Println("error in creating ", err)
	if err != nil {
		return err
	}
	return nil
}
