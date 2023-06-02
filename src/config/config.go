package config

import (
	"os"
	"log"
	"fmt"
	"github.com/joho/godotenv"
	"strconv"
)


type DBConfig struct {
	Host 		string
	Port 		int
	User 		string
	DBName 		string
	Password	string
}


var Configs []DBConfig

func GetConfig() []DBConfig {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading config file")
	}

	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_USER = os.Getenv("DB_USER")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")

	PORT, _ := strconv.Atoi(DB_PORT)

	Configs = append(Configs, DBConfig {
		Host: DB_HOST, 
		Port: PORT, 
		User: DB_USER, 
		DBName: DB_NAME, 
		Password: DB_PASSWORD,
	})

	fmt.Println(Configs)
	return Configs
   
}