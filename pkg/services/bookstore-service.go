package services

import (
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

var BookstoreInterface domain.IBookstoreInterface

func SetBookstoreInterface(bInterface domain.IBookstoreInterface) {
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
func Get(bookID uint) types.BookRequest {
	book := BookstoreInterface.Get(bookID)
	return types.BookRequest{
		BookName:    book.BookName,
		Author:      book.Author,
		Publication: book.Publication,
	}
}
func Create(book *models.Book) error {
	if err := BookstoreInterface.Create(book); err != nil {
		return err
	}
	return nil
}

func Update(book *models.Book) error {
	if err := BookstoreInterface.Update(book); err != nil {
		return err
	}
	return nil
}

func Delete(bookID uint) error {
	if err := BookstoreInterface.Delete(bookID); err != nil {
		return err
	}
	return nil
}
