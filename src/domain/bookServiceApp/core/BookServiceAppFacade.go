package core

import (
	"book-service-app/src/domain/bookServiceApp/core/ports/incoming"
	"book-service-app/src/domain/bookServiceApp/core/model"
)

type BookService struct {
    bookService incoming.BookRepository
}

func NewBookService (bookService incoming.BookRepository) *BookService {
	return &BookService{bookService}
}

func (b *BookService) GetBooks() ([]model.Book, error) {
	return b.bookService.GetAllBooks() 
}

func (b *BookService) GetBookByID(id string) (*model.Book, error) {
	return b.bookService.GetByID(id)
}

func (b *BookService) CreateBooks(book model.Book) ( error) {
	return b.bookService.CreateBook(book)
}

func (b *BookService) UpdateBookByID(book model.Book) (*model.Book, error) {
	return b.bookService.UpdateBook(book)
}

func (b *BookService) DeleteBookByID(id string) (error) {
	return b.bookService.DeleteBook(id)
}






