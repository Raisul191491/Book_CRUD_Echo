package domain

import (
	"go-bootcamp/pkg/infra/connection/db/model"
	"go-bootcamp/pkg/types"
)

type IBookstoreRepo interface {
	GetBooks(bookID uint) []model.Book
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(bookID uint) error
}

type IBookstoreService interface {
	GetBooks(bookID uint) ([]types.BookRequest, error)
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(bookID uint) error
}
