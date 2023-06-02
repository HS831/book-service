package routes

import (
	"github.com/gin-gonic/gin"
	store "book-service-app/src/domain/bookServiceApp/infrastructure"
	"book-service-app/src/domain/bookServiceApp/core"
	"book-service-app/src/domain/bookServiceApp/infrastructure/adapters"
	db "book-service-app/src/store"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func Run() {
	router := gin.Default()

	// intialized book store
	bookStore := store.NewBookStore(db.DB)

	// initialized book service
	bookService := core.NewBookService(bookStore)

	// initialized service adapter
	serviceAdapter := adapters.NewServiceAdapter(bookService)

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", serviceAdapter.GetAllBooksHandler)
	router.GET("/books/:id", serviceAdapter.GetBooksByIDHandler)
	router.POST("/books", serviceAdapter.CreateBooksHandler)
	router.PATCH("/books/:id", serviceAdapter.UpdateBooksByIDHandler)
	router.DELETE("books/:id", serviceAdapter.DeleteBooksByIDHandler)

	router.Run("localhost:3000")
}