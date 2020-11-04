package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"time"

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
	logPath  string
	Instance *gorm.DB
}

var instance *Database
var once sync.Once

//Setup instantiates DB, opens the DB connection and sets up DB schema.
func Setup(DBName, Location, User, Password, Port, logPath string) {
	database := getInstance(DBName, Location, User, Password, Port, logPath)
	database.open()
}

//Shut closes the input DB connection.
func Shut(DBName, Location, User, Password, Port, logPath string) {
	database := getInstance(DBName, Location, User, Password, Port, logPath)
	database.Close()
}

func getInstance(DBName, Location, User, Password, Port, logPath string) *Database {
	once.Do(func() {
		instance = &Database{Name: DBName, Location: Location, User: User, Password: Password, Port: Port, logPath: logPath}
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
	logPath := database.logPath

	newLogger := database.SetupDBLogger(logPath)

	url := user + ":" + password + "@tcp(" + location + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True"
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	if err != nil {
		msg := fmt.Errorf("Error in opening connection to database with URL %s, Reason: %v ", url, err)
		fmt.Println(msg)
	}
	sqlDB, _ := conn.DB()
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(10)
	fmt.Println("Connection Established")
	database.Instance = conn
	return err
}

func (database *Database) SetupDBLogger(logPath string) logger.Interface {
	var file io.Writer
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		file, _ = os.Create(logPath)
	} else {
		file, _ = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	}
	newLogger := getDBLogger(file)
	return newLogger
}

func getDBLogger(file io.Writer) logger.Interface {
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	return newLogger
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
