package services

import (
	"errors"
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
)

type BookService struct {
	bService domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		bService: bookRepo,
	}
}

func (service *BookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.bService.GetBooks(bookID)
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

func (service *BookService) CreateBook(book *models.Book) error {
	if err := service.bService.CreateBook(book); err != nil {
		return errors.New("Book was not created")
	}
	return nil
}

func (service *BookService) UpdateBook(book *models.Book) error {
	if err := service.bService.UpdateBook(book); err != nil {
		return errors.New("Book update was unsuccesful")
	}
	return nil
}

func (service *BookService) DeleteBook(bookID uint) error {
	if err := service.bService.DeleteBook(bookID); err != nil {
		return errors.New("Book deletion was unsuccesful")
	}
	return nil
}
