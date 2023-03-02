package services

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

var BookstoreInterface domain.IBookstoreRepo

type BookService struct {
	bService domain.IBookstoreRepo
}

func BookStoreServiceInstance(bookRepo domain.IBookstoreRepo) domain.IBookstoreService {
	return &BookService{
		bService: bookRepo,
	}
}

func SetBookstoreInterface(bInterface domain.IBookstoreRepo) {
	BookstoreInterface = bInterface
}

func BookList() []types.CustomBookResponse {
	var allBooks []types.CustomBookResponse
	bookList := BookstoreInterface.BookList()
	for _, val := range bookList {
		allBooks = append(allBooks, types.CustomBookResponse{
			ID:       val.ID,
			BookName: val.BookName,
		})
	}
	return allBooks
}
func (b *BookService) GetBooks(bookID uint) types.BookRequest {
	book := BookstoreInterface.GetBooks(bookID)
	return types.BookRequest{
		BookName:    book.BookName,
		Author:      book.Author,
		Publication: book.Publication,
	}
}
func (b *BookService) CreateBook(book *models.Book) error {
	if err := BookstoreInterface.CreateBook(book); err != nil {
		return err
	}
	return nil
}

func (b *BookService) UpdateBook(book *models.Book) error {
	if err := BookstoreInterface.UpdateBook(book); err != nil {
		return err
	}
	return nil
}

func (b *BookService) DeleteBook(bookID uint) error {
	if err := BookstoreInterface.DeleteBook(bookID); err != nil {
		return err
	}
	return nil
}
