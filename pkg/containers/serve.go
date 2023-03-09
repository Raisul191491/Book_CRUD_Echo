package containers

import (
	"go-bootcamp/pkg/connection"
	"go-bootcamp/pkg/controllers"
	"go-bootcamp/pkg/repositories"
	"go-bootcamp/pkg/routes"
	"go-bootcamp/pkg/services"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	connection.Connect()
	db := connection.GetDB()

	bookRepo := repositories.BookDBInstance(db)
	services.SetBookInterface(bookRepo)

	bookService := services.BookServiceInstance(bookRepo)
	controllers.SetBookService(bookService)

	routes.BookRoutes(e)
	log.Fatal(e.Start(":9030"))
}
