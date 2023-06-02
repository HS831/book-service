package store

import (
	"fmt"
	"go.uber.org/zap"
	config "book-service-app/src/config"
	"book-service-app/src/domain/bookServiceApp/core/model"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	log "book-service-app/src/utils/loggerUtils"
)

var DB *gorm.DB

type DBConfig = config.DBConfig

func BuildDB () *DBConfig {
	// err := godotenv.Load("../.env")

	// if err != nil {
	// 	log.Panic("Error loading config file")
	// }

	// var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	// var DB_NAME = os.Getenv("DB_NAME")
	// var DB_USER = os.Getenv("DB_USER")
	// var DB_HOST = os.Getenv("DB_HOST")
	// var DB_PORT = os.Getenv("DB_PORT")

	// PORT, _ := strconv.Atoi(DB_PORT)
	
	dbConfig := DBConfig {
		Host: "localhost", 
		Port: 3306, 
		User: "root", 
		DBName: "bookService", 
		Password: "harry@831just",
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

var err error
func ConnectDB() {
	logger := log.FileLogger()
	DB, err = gorm.Open("mysql", DbURL(BuildDB())) 
	if err != nil {
		logger.Error("Status: ", zap.Error(err))
	} else {
		logger.Info("Connection Established")
	}

	DB.AutoMigrate(&model.Book{})
}

