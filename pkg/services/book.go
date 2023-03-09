package services

import (
	"errors"
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

var BookInterface domain.IBookRepo

type BookService struct {
	bService domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		bService: bookRepo,
	}
}

func SetBookInterface(bInterface domain.IBookRepo) {
	BookInterface = bInterface
}

func (b *BookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := BookInterface.GetBooks(bookID)
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
	if err := BookInterface.CreateBook(book); err != nil {
		return errors.New("Book was not created")
	}
	return nil
}

func (b *BookService) UpdateBook(book *models.Book) error {
	if err := BookInterface.UpdateBook(book); err != nil {
		return errors.New("Book update was unsuccesful")
	}
	return nil
}

func (b *BookService) DeleteBook(bookID uint) error {
	if err := BookInterface.DeleteBook(bookID); err != nil {
		return errors.New("Book deletion was unsuccesful")
	}
	return nil
}
