package routes

import (
	"go-bootcamp/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func BookStoreRoutes(e *echo.Echo) {
	bookstore := e.Group("/bookstore")

	bookstore.POST("/book", controllers.CreateBook)
	bookstore.GET("/book/:bookID", controllers.GetBooks)
	bookstore.GET("/booklist", controllers.GetBookList)
	bookstore.PUT("/book/:bookID", controllers.UpdateBook)
	bookstore.DELETE("/book/:bookID", controllers.DeleteBook)
}
