package outgoing

import (
	book "book-service-app/src/domain/bookServiceApp/core/model"
)
type Books = book.Book

type BookStore interface {
	GetAllBooks() ([]Books, error)
	GetByID(id string) (*Books, error)
	CreateBook(book Books) error
	UpdateBook(book Books ) (*Books, error)
	DeleteBook(id string ) error
}