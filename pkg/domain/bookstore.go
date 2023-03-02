package domain

import (
	"go-bootcamp/pkg/models"
)

type IBookstoreInterface interface {
	BookList() []models.Book
	GetBooks(bookID uint) models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}
