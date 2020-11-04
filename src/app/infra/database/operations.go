package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	//mysql dialects
	"gorm.io/driver/mysql"
	"os"
	"sync"
)

//Database stores the application data.
type Database struct {
	Name     string
	Location string
	User     string
	Password string
	Port     string
	Instance *gorm.DB
}

var instance *Database
var once sync.Once

//Setup instantiates DB, opens the DB connection and sets up DB schema.
func Setup(DBName, Location, User, Password, Port string) {
	database := getInstance(DBName, Location, User, Password, Port)
	database.open()
}

//Shut closes the input DB connection.
func Shut(DBName, Location, User, Password, Port string) {
	database := getInstance(DBName, Location, User, Password, Port)
	database.Close()
}

func getInstance(DBName, Location, User, Password, Port string) *Database {
	once.Do(func() {
		instance = &Database{Name: DBName, Location: Location, User: User, Password: Password, Port: Port}
	})
	return instance
}

//GetWorkingInstance returns the current working instance of the DB.
//should be called only after database has been setup
func GetWorkingInstance() *Database {
	return instance
}

func (database *Database) open() (err error) {
	dbName := database.Name
	user := database.User
	password := database.Password
	location := database.Location
	port := database.Port
	url := user + ":" + password + "@tcp(" + location + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True"
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := database.Instance.DB()
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(10)
	fmt.Println("Connection Established")
	database.Instance = conn
	return err
}

//MigrateSchema Migrate DB schema for unit / api integration test code
func (database *Database) MigrateSchema() (err error) {
	database.Instance.AutoMigrate(&App{})
	return nil
}

//Close closes the DB connection.
func (database *Database) Close() (err error) {
	sqlDB, _ := database.Instance.DB()
	return sqlDB.Close()
}

func (database *Database) delete() (err error) {
	return os.Remove(database.Name)
}
