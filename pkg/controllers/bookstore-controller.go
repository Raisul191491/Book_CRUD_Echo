package controllers

import (
	"go-bootcamp/pkg/consts"
	"go-bootcamp/pkg/models"
	"go-bootcamp/pkg/services"
	"go-bootcamp/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBookList(e echo.Context) error {
	bookList := services.BookList()
	if len(bookList) == 0 {
		return e.JSON(http.StatusOK, consts.InvalidInput)
	}
	return e.JSON(http.StatusOK, bookList)
}

func GetBooks(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	book := services.Get(uint(bookID))

	if book.BookName == "" || book.Author == "" {
		return e.JSON(http.StatusBadRequest, consts.NotFound)
	}

	return e.JSON(http.StatusOK, book)
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

	if err := services.Create(book); err != nil {
		return e.JSON(http.StatusInternalServerError, consts.Failure)
	}

	return e.JSON(http.StatusCreated, consts.Success)
}
func UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}

	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}

	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	book := services.Get(uint(bookID))

	if book.BookName == "" || book.Author == "" {
		return e.JSON(http.StatusBadRequest, consts.NotFound)
	}

	updatedBook := ChangeBookInfo(*reqBook)
	updatedBook.ID = uint(bookID)

	if err := services.Update(&updatedBook); err != nil {
		return e.JSON(http.StatusInternalServerError, consts.Failure)
	}

	return e.JSON(http.StatusCreated, consts.Success)
}
func DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseUint(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, consts.InvalidID)
	}
	book := services.Get(uint(bookID))

	if book.BookName == "" || book.Author == "" {
		return e.JSON(http.StatusBadRequest, consts.NotFound)
	}

	if err := services.Delete(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, consts.Failure)
	}

	return e.JSON(http.StatusOK, consts.Success)
}

func ChangeBookInfo(reqBook types.BookRequest) models.Book {
	var updatedBook models.Book

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
