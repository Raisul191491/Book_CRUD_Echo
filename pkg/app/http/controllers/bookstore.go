package controllers

import (
	"go-bootcamp/pkg/app/domain"
	"go-bootcamp/pkg/app/utils/consts"
	"go-bootcamp/pkg/infra/connection/db/model"

	"go-bootcamp/pkg/app/utils/types"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var BookstoreService domain.IBookstoreService

func SetBookstoreService(bService domain.IBookstoreService) {
	BookstoreService = bService
}

func CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}

	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidInput)
	}

	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	book := &model.Book{
		BookName:    reqBook.BookName,
		Author:      reqBook.Author,
		Publication: reqBook.Publication,
	}

	if err := BookstoreService.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Book was created successfully")
}

func GetBooks(e echo.Context) error {
	tempBookID := e.QueryParam("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil && tempBookID != "" {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	book, err := BookstoreService.GetBooks(uint(bookID))

	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, book)
}

func UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}

	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidInput)
	}

	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	_, err = BookstoreService.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook := ChangeBookInfo(*reqBook)
	updatedBook.ID = uint(bookID)

	if err := BookstoreService.UpdateBook(&updatedBook); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Book was updated successfully")
}
func DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	_, err = BookstoreService.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := BookstoreService.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "Book was deleted successfully")
}

func ChangeBookInfo(reqBook types.BookRequest) model.Book {
	var updatedBook model.Book

	if reqBook.BookName != "" {
		updatedBook.BookName = reqBook.BookName
	}
	if reqBook.Author != "" {
		updatedBook.Author = reqBook.Author
	}
	if reqBook.Publication != "" {
		updatedBook.Publication = reqBook.Publication
	}

	return updatedBook
}
