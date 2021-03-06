package usecase

import (
	"fmt"
	"go-be-book/books"
	"go-be-book/models"

	"gorm.io/gorm"
)

type bookUsecase struct {
	booksRepository        books.Repository
}

// NewCompanyUsecase is constructor for company usecase.
func NewBooksUsecase(
	booksRepository        books.Repository,
) books.Usecase {
	return &bookUsecase{
		booksRepository,
	}
}

// ListCompanies return company list.
func (usecase *bookUsecase) ListBooks(db *gorm.DB) ([]models.Book, error) {
	books, err := usecase.booksRepository.ListBook(db)

	switch err {
	case nil:
		return books, nil
	default:
		fmt.Println("Get list Books failed !")
		return nil, err
	}
}
