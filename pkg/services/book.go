package services

import (
	"errors"
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

func (b *BookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := BookstoreInterface.GetBooks(bookID)
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			ID:          val.ID,
			BookName:    val.BookName,
			Author:      val.Author,
			Publication: val.Publication,
		})
	}
	if len(book) == 0 {
		return nil, errors.New("No book found")
	}
	return allBooks, nil
}

func (b *BookService) CreateBook(book *models.Book) error {
	if err := BookstoreInterface.CreateBook(book); err != nil {
		return errors.New("Book was not created")
	}
	return nil
}

func (b *BookService) UpdateBook(book *models.Book) error {
	if err := BookstoreInterface.UpdateBook(book); err != nil {
		return errors.New("Book update was unsuccesful")
	}
	return nil
}

func (b *BookService) DeleteBook(bookID uint) error {
	if err := BookstoreInterface.DeleteBook(bookID); err != nil {
		return errors.New("Book deletion was unsuccesful")
	}
	return nil
}
