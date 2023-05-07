package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
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


func main () {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.PATCH("/books/:id", patchBooks)
	router.DELETE("books/:id", deleteBookById)

	router.Run("localhost:8000")
}

func getAllBooks (r *gin.Context) {
	r.IndentedJSON(http.StatusOK, books)
}

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

func postBooks (r *gin.Context) {
	var newBook book
	
	if err := r.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	r.IndentedJSON(http.StatusCreated, newBook)
}

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