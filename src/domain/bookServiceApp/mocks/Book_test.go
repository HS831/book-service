package mocks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"book-service-app/src/domain/bookServiceApp/core/model"
	"book-service-app/src/domain/bookServiceApp/core"
	"book-service-app/src/domain/bookServiceApp/infrastructure/adapters"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/jinzhu/gorm"
)

var mockDB *MockDB

func TestBookAPI(t *testing.T) {
	// Set up the mock database
	mockDB = new(MockDB)
	router := gin.Default()

	// intialized book store
	bookStore := NewMockBookStore(mockDB)

	// initialized book service
	bookService := core.NewBookService(bookStore)

	// initialized service adapter
	serviceAdapter := adapters.NewServiceAdapter(bookService)

	router.GET("/books", serviceAdapter.GetAllBooksHandler)
	router.GET("/books/:id", serviceAdapter.GetBooksByIDHandler)
	router.POST("/books", serviceAdapter.CreateBooksHandler)
	router.PATCH("/books/:id", serviceAdapter.UpdateBooksByIDHandler)
	router.DELETE("books/:id", serviceAdapter.DeleteBooksByIDHandler)
	
	// Test Create (POST /books)
	t.Run("CreateBook", func(t *testing.T) {
		newBook := model.Book{ID: "1", Title: "Book 1", Author: "Author 1"}
		payload, _ := json.Marshal(newBook)

		mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

		req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		var createdBook model.Book
        err = json.Unmarshal(recorder.Body.Bytes(), &createdBook)
        assert.NoError(t, err)

        assert.Equal(t, newBook.Title, createdBook.Title)
        assert.Equal(t, newBook.Author, createdBook.Author)
	})

	// Test Read (GET /books)
	t.Run("GetBooks", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/books", nil)
		assert.NoError(t, err)

		mockDB.On("Find", mock.Anything).Return(&gorm.DB{})

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		
		assert.Equal(t, http.StatusOK, recorder.Code)

		var books []model.Book
		err = json.Unmarshal(recorder.Body.Bytes(), &books)
		assert.NoError(t, err)

		assert.NotEmpty(t, books)
	})

	// Test Read (GET /books/:id)
	t.Run("GetBookByID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/books/1", nil)
		assert.NoError(t, err)

		mockDB.On("First", mock.Anything, "1").Return(&gorm.DB{})

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var retrievedBook model.Book
		err = json.Unmarshal(recorder.Body.Bytes(), &retrievedBook)
		assert.NoError(t, err)

		assert.Equal(t, "Book 1", retrievedBook.Title)
		assert.Equal(t, "Author 1", retrievedBook.Author)
	})

	// Test Update (PATCH /books/:id)
	t.Run("UpdateBook", func(t *testing.T) {
		updatedBook := model.Book{ID : "1", Title: "Updated Book", Author: "Updated Author"}
		payload, _ := json.Marshal(updatedBook)

		mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

		req, err := http.NewRequest("PATCH", "/books/1", bytes.NewBuffer(payload))
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		
		var responseBook model.Book
		err = json.Unmarshal(recorder.Body.Bytes(), &responseBook)
		assert.NoError(t, err)

		assert.Equal(t, updatedBook.Title, responseBook.Title)
		assert.Equal(t, updatedBook.Author, responseBook.Author)
		assert.Equal(t, "1", responseBook.ID)
	})

	// Test Delete (DELETE /books/:id)
	t.Run("DeleteBook", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/books/1", nil)
		assert.NoError(t, err)

		mockDB.On("Delete", mock.Anything).Return(&gorm.DB{})

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Book with given ID deleted", response["message"])
	})
}