package domain

import (
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

type IBookstoreRepo interface {
	GetBooks(bookID uint) []models.Book
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}

type IBookstoreService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(bookID uint) error
}
