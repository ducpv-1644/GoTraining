package books

import (
	"go-be-book/models"

	"gorm.io/gorm"
)

type Repository interface {
	ListBook(db *gorm.DB) ([]models.Book, error)
	CreateBook(db *gorm.DB, title string, author string) (models.Book, error)
	RetrieveBook(db *gorm.DB, id string) (models.Book, error)
	DeleteBook(db *gorm.DB, id string)
	UpdateBook(db *gorm.DB, id string, title string, author string) (models.Book)
	CreateBookWithChannels(db *gorm.DB, bookTitle string, bookAuthor string, chlBook chan models.Book)
}
