package main

import (
	_ "github.com/go-sql-driver/mysql"
	db "book-service-app/src/store"
	"book-service-app/src/domain/bookServiceApp/application/routes"
	"book-service-app/src/utils"
	log "book-service-app/src/utils/loggerUtils"
)

// @Title 			Book Service API
// @Version			v1.0.0
// @Description		A Book-Service API which contains all the basic CRUD functionality for any book-service application using Go and Gin Framework.

// @Host			localhost:3000
// @BasePath		/

func main () {

	// Intialize loggers
	logger := log.FileLogger()
	logger.Error("Getting some error")

	// Create a code management file
	utils.CodeManager()

	// Establish connection to Database.
	db.ConnectDB()

	// Start the Application
	routes.Run()
}
