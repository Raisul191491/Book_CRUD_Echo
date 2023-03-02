package domain

import (
	"go-bootcamp/pkg/models"
)

type IBookstoreInterface interface {
	BookList() []models.Book
	Get(bookID uint) models.Book
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(bookID uint) error
}
