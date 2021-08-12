package repository

import (
	"go-be-book/models"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

type BookRepository struct{}

func (*BookRepository) ListBook(db *gorm.DB) ([]models.Book, error) {
	books := []models.Book{}
	db.Find(&books)
	return books, nil
}

func (*BookRepository) CreateBook(db *gorm.DB, title string, author string) (models.Book, error) {
	book := models.Book{
		Title: title,
		Author: author,
	}
	db.Create(&book)
	return book, nil
}

func (*BookRepository) RetrieveBook(db *gorm.DB, id string) (models.Book, error) {

	book := models.Book{}
	db.First(&book, id)
	return book, nil
}

func (*BookRepository) DeleteBook(db *gorm.DB, id string) (string, error) {

	db.Delete(&models.Book{}, id)
	return id, nil
}
