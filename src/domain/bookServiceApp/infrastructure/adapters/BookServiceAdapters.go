package adapters

import (
    "github.com/gin-gonic/gin"
    "book-service-app/src/domain/bookServiceApp/core"
    "book-service-app/src/domain/bookServiceApp/core/model"
    "net/http"
    _ "book-service-app/src/docs"
)

type ServiceAdapter struct {
    bookService *core.BookService
}


func NewServiceAdapter(bookService *core.BookService) *ServiceAdapter {
    return &ServiceAdapter{bookService}
}


// GetAllBooks 		is a swagger annotation for the GET request to fetch all the books from the server.
// @Summary 		Get all books.
// @Description 	This endpoint fetches all the books available on the server.
// @Tags            books
// @Produce 		json
// @Tags 			books
// @Success 		200 {array} book "Successful operation"
// @Failure 		404 {object} gin.H
// @Failure 		500 {object} gin.H
// @Router 			/books [get]

func (b *ServiceAdapter) GetAllBooksHandler(c *gin.Context) {

    books, err := b.bookService.GetBooks()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, books)
}

// PostBooks 		is a swagger annotation for the POST request to add a new book to the server.
// @Summary 		Add a new book.
// @Description 	This endpoint adds a new book to the server.
// @Accept 		    json
// @Produce 		json
// @Tags 			books
// @Param 			newBook body book true "New book details"
// @Success 		201 {object} book "Successful operation"
// @Failure 		400 {object} gin.H
// @Failure 		500 {object} gin.H
// @Router 			/books [post]

func (b *ServiceAdapter) CreateBooksHandler(c *gin.Context) {
    var newBook model.Book
	c.BindJSON(&newBook)
    err := b.bookService.CreateBooks(newBook)
   
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newBook)
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

func (b *ServiceAdapter) GetBooksByIDHandler(c *gin.Context) {
    id := c.Param("id")
    book, err := b.bookService.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Book with given ID not present in the database."})
        return
    }
    c.JSON(http.StatusOK, book)
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

func (b *ServiceAdapter) UpdateBooksByIDHandler(c *gin.Context) {
    id := c.Param("id")
    _, err := b.bookService.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message" : "Book with given ID not present in the database."})
    }

    var updatedBook model.Book
    c.BindJSON(&updatedBook)

    book, updErr := b.bookService.UpdateBookByID(updatedBook)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": updErr.Error()})
        return
    }

    c.JSON(http.StatusOK, book)
}

// @Summary 		Delete an existing book.
// @Description 	Delete the book with the ID that matches the parameter.
// @Param 			id path string true "The ID of the book to delete"
// @Produce 		json
// @Tags 			books
// @Success 		200 {object} gin.H "Successful operation"
// @Failure 		404 {object} gin.H
// @Router 			/books/{id} [delete]

func (b *ServiceAdapter) DeleteBooksByIDHandler(c *gin.Context) {
    id := c.Param("id")
    _ , err := b.bookService.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message" : "Book with given ID not present in the database."})
    }

    delErr := b.bookService.DeleteBookByID(id)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": delErr.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book with given ID deleted"})
}