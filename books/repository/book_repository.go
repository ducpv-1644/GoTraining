package repository

import (
	"fmt"
	"time"
	"go-be-book/books"
	"go-be-book/models"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

func NewBookRepository() books.Repository {
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

func (*BookRepository) DeleteBook(db *gorm.DB, id string) {
	db.Delete(&models.Book{}, id)
}

func (*BookRepository) UpdateBook(db *gorm.DB, id string, title string, author string) (models.Book){
	book := models.Book{}
	db.First(&book, id)
	book.Author = author
	book.Title = title
	db.Save(&book)
	return book
}

func (*BookRepository) CreateBookWithChannels(db *gorm.DB, bookTitle string, bookAuthor string, chlBook chan models.Book) {
	book := models.Book{
		Title: bookTitle,
		Author: bookAuthor,
	}
	db.Create(&book)

	tmn := time.Now()
	fmt.Println("")
	fmt.Println(tmn)
	fmt.Println(book)
	fmt.Println("")

	chlBook <- book
}
