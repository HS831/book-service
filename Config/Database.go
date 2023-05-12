package Config

import (
	"fmt"
	"os"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

type DBConfig struct {
	Host 		string
	Port 		uint
	User 		string
	DBName 		string
	Password	string
}

func BuildDB () *DBConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var DB_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	var DB_NAME = os.Getenv("DATABASE_NAME")

	dbConfig := DBConfig {
		Host: "localhost",
		Port: 3306,
		User: "root",
		Password: DB_PASSWORD,
		DBName: DB_NAME,
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
