package mocks

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/jinzhu/gorm"
	"book-service-app/src/domain/bookServiceApp/core/model"
)


type MockDB struct {
	mock.Mock
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(dest interface{}) *gorm.DB {
	args := m.Called(dest)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, id string) *gorm.DB {
	args := m.Called(dest, id)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}


type MockBookStore struct {
    db *MockDB
}

func NewMockBookStore(db *MockDB) *MockBookStore {
    return &MockBookStore{db: db}
}

var data map[string] model.Book

func makeMap() {
	data = make(map[string] model.Book)
}

func (m *MockBookStore) CreateBook(book model.Book) error {
	fmt.Println(book.ID)
	makeMap()
    result := m.db.Create(book)
	data[book.ID] = book
    return result.Error
}

func (m *MockBookStore) GetAllBooks() ([]model.Book, error) {
    len := len(data)
	if len == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var books []model.Book
	for _, book := range data {
		books = append(books, book)
	}

	return books, nil
}

func (m *MockBookStore) GetByID(id string) (*model.Book, error) {
	book := data[id]
	return &book, nil
}

func (m *MockBookStore) UpdateBook(book model.Book) (*model.Book, error) {
    result := m.db.Save(book)
    return &book, result.Error
}

func (m *MockBookStore) DeleteBook(id string) error {
    book := &model.Book{ID: id}
    result := m.db.Delete(book)
    return result.Error
}