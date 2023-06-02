package infrastructure


import (
    "time"
    "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"book-service-app/src/domain/bookServiceApp/core/model"
)

type BookStore struct {
    db *gorm.DB
}

func NewBookStore(db *gorm.DB) *BookStore {
    return &BookStore{db}
}

func (b *BookStore) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
    maxRetries := 3
    retryDelay := 1 * time.Second
    var err error

    for i:= 0; i<maxRetries; i++ {
        err = b.db.Find(&books).Error
        
        if err == nil {
            return books, nil
        } 

        fmt.Println("Error retrieving books: Retrying in ....", err, retryDelay)
    }
    return nil, err
}

func (b *BookStore) GetByID(id string) (*model.Book, error) {
    var book model.Book
    result := b.db.First(&book, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &book, nil
}

func (b *BookStore) CreateBook(newbook model.Book) (error) {
    result := b.db.Create(newbook)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (b *BookStore) UpdateBook(updatedBook model.Book) (*model.Book, error) {

    result := b.db.Save(updatedBook)
    if result.Error != nil {
        return nil, result.Error
    }

    return &updatedBook, nil
}

func (b *BookStore) DeleteBook(id string) (error) {
    var delBook model.Book
    result := b.db.Where("ID = ?", id).Delete(&delBook)
    if result.Error != nil {
        return result.Error
    }

    return nil
}