package controllers

import (
	"go-bootcamp/pkg/consts"
	"go-bootcamp/pkg/domain"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var BookService domain.IBookService

func SetBookService(bService domain.IBookService) {
	BookService = bService
}

func CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}

	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidInput)
	}

	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	book := &models.Book{
		BookName:    reqBook.BookName,
		Author:      reqBook.Author,
		Publication: reqBook.Publication,
	}

	if err := BookService.CreateBook(book); err != nil {
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

	book, err := BookService.GetBooks(uint(bookID))

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

	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	existingBook, err := BookService.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedBook := ChangeBookInfo(existingBook[0], *reqBook)
	updatedBook.ID = uint(bookID)

	checkBook := &types.BookRequest{
		BookName:    updatedBook.BookName,
		Author:      updatedBook.Author,
		Publication: updatedBook.Publication,
	}

	if err := checkBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := BookService.UpdateBook(&updatedBook); err != nil {
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

	_, err = BookService.GetBooks(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := BookService.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "Book was deleted successfully")
}

func ChangeBookInfo(oldBook, reqBook types.BookRequest) models.Book {

	var updatedBook models.Book
	updatedBook.BookName = oldBook.BookName
	updatedBook.Author = oldBook.Author
	updatedBook.Publication = oldBook.Publication

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
