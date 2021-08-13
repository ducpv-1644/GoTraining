package books

import (
	"go-be-book/models"

	"gorm.io/gorm"
)

type Usecase interface {
	ListBooks(db *gorm.DB) ([]models.Book, error)
}
