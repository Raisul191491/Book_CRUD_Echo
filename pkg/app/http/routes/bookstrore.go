package routes

import (
	"go-bootcamp/pkg/app/http/controllers"

	"github.com/labstack/echo/v4"
)

func BookStoreRoutes(e *echo.Echo) {
	bookstore := e.Group("/bookstore")

	bookstore.POST("/book", controllers.CreateBook)
	bookstore.GET("/book", controllers.GetBooks)
	bookstore.PUT("/book/:bookID", controllers.UpdateBook)
	bookstore.DELETE("/book/:bookID", controllers.DeleteBook)
}
