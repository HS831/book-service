package main

import (
	"fmt"
	"net/http"
	_ "main/docs"

	"main/Config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type book struct {
	ID 		string 
	Title 	string
	Author 	string			
	Price 	string			
}

// var books = []book {
// 	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Price: "$399.90"},
// 	{ID: "2", Title: "The Castle", Author: "Franz Kafka", Price: "$145.98"},
// }
var err error

func connectDB() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDB())) 
	if err != nil {
		fmt.Println("Status: ", err)
	}

	Config.DB.AutoMigrate(&book{})

}

// @title 			Book Service API
// @version			v1.0.0
// @description		A Book-Service API which contains all the basic CRUD functionality for any book-service application using Go and Gin Framework.

// @host			localhost:3000

func main () {
	connectDB()

	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.PUT("/books/:id", patchBooks)
	router.DELETE("/books/:id", deleteBookById)

	router.Run("localhost:3000")

	defer Config.DB.Close()
}

// GetAllBooks 		is a swagger annotation for the GET request to fetch all the books from the server.
// @Summary 		Get all books.
// @Description 	This endpoint fetches all the books available on the server.
// @Produce 		json
// @Tags 			books
// @Success 		200 {array} book "Successful operation"
// @Failure 		404 {object} gin.H
// @Failure 		500 {object} gin.H
// @Router 			/books [get]


func getAllBooks (r *gin.Context) {
	var books []book
	err = Config.DB.Find(&books).Error; 
	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
	} else {
		r.JSON(http.StatusOK, books)
	}
}


// GetBookByID 		is a swagger annotation for the GET request to fetch a single book by its ID from the server.
// @Summary 		Get a single book.
// @Description 	This endpoint fetches the book with the ID that matches the parameter.
// @Param 			id path string true "The ID of the book to fetch"
// @Produce 		json
// @Tags 			books
// @Success 		200 {object} book "Successful operation"
// @Failure 		404 {object} gin.H
// @Failure 		500 {object} gin.H
// @Router 			/books/{id} [get]


func getBookById (r *gin.Context) {
	id := r.Param("id")

	var oneBook book
	err := Config.DB.Where("ID = ?", id).First(&oneBook).Error

	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
	} else {
		r.JSON(http.StatusOK, oneBook)
	}
}

// PostBooks 		is a swagger annotation for the POST request to add a new book to the server.
// @Summary 		Add a new book.
// @Description 	This endpoint adds a new book to the server.
// @Accept 		json
// @Produce 		json
// @Tags 			books
// @Param 			newBook body book true "New book details"
// @Success 		201 {object} book "Successful operation"
// @Failure 		400 {object} gin.H
// @Failure 		500 {object} gin.H
// @Router 			/books [post]

func postBooks (r *gin.Context) {
	var newBook book
	r.BindJSON(&newBook)
	err := Config.DB.Create(newBook).Error

	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
	} else {
		r.JSON(http.StatusOK, newBook)
	}
}

// @Summary 		Update an existing book.
// @Description 	Update the book with the ID that matches the parameter.
// @Param 			id path string true "The ID of the book to update"
// @Param 			book body book true "The updated book"
// @Produce 		json
// @Tags 			books
// @Success 		200 {object} book "Successful operation"
// @Failure 		404 {object} gin.H
// @Router 			/books/{id} [patch]

func patchBooks (r *gin.Context) {
	//id := r.Param("id")

    var newBook book
	r.BindJSON(&newBook)
	//err := nil

	
		Config.DB.Save(&newBook)

		r.JSON(http.StatusOK, newBook)
	
}

// @Summary 		Delete an existing book.
// @Description 	Delete the book with the ID that matches the parameter.
// @Param 			id path string true "The ID of the book to delete"
// @Produce 		json
// @Tags 			books
// @Success 		200 {object} gin.H "Successful operation"
// @Failure 		404 {object} gin.H
// @Router 			/books/{id} [delete]

func deleteBookById (r *gin.Context) {
	id := r.Param("id")
	var delBook book

	Config.DB.Where("ID = ?", id).Delete(&delBook)

	r.JSON(http.StatusOK, gin.H{"message" : "Book with given ID deleted!!"})
}

