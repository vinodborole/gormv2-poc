package database

import "gorm.io/gorm"

//App struct
type App struct {
	gorm.Model

	Name        string
	Description string
	Url         string
	Port        string
}
