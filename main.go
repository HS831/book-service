package main

import (
	"net/http"
	_ "main/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type book struct {
	ID string
	Title string
	Author string
	Price string
}

var books = []book {
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Price: "$399.90"},
	{ID: "2", Title: "The Castle", Author: "Franz Kafka", Price: "$145.98"},
}

// @title 			Book Service API
// @version			v1.0.0
// @description		A Book-Service API which contains all the basic CRUD functionality for any book-service application using Go and Gin Framework.

// @host			localhost:3000
// @BasePath 		/api
// @Tags			/books
func main () {
	router := gin.Default()

	// adding swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.PATCH("/books/:id", patchBooks)
	router.DELETE("books/:id", deleteBookById)

	router.Run("localhost:3000")
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
// @OperationID 	getAllBooks


func getAllBooks (r *gin.Context) {
	r.IndentedJSON(http.StatusOK, books)
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
// @OperationID     getBookById

func getBookById (r *gin.Context) {
	id := r.Param("id")

	for _ , b := range books {
		if(b.ID == id) {
			r.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	r.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book with given ID not found!!"})
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
// @OperationID 	createBook

func postBooks (r *gin.Context) {
	var newBook book
	
	if err := r.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	r.IndentedJSON(http.StatusCreated, newBook)
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
	id := r.Param("id")

    var newBook book
	
    if err := r.BindJSON(&newBook); err!= nil {
        return
    }

	for _ , b := range books {
		if (b.ID == id) {
			b.Title = newBook.Title
			b.Author = newBook.Author
			b.Price = newBook.Price
			r.IndentedJSON(http.StatusOK, b)
            return
		}
	}

	r.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book with given ID not found!!"})
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
	var delIndex = -1

	for index, b := range books {
		if(b.ID == id) {
			delIndex = index
			break;
		}
	}

	if(delIndex == -1) {
		r.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book with given ID not found!!"})
        return
	}

	books = append(books[:delIndex], books[delIndex+1:]...)

	r.IndentedJSON(http.StatusOK, gin.H{"message" : "Book with given ID deleted!!"})
}