package containers

import (
	"go-bootcamp/pkg/connection"
	"go-bootcamp/pkg/repositories"
	"go-bootcamp/pkg/routes"
	"go-bootcamp/pkg/services"
	"log"

	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo) {
	connection.Connect()
	db := connection.GetDB()

	bookstoreInterface := repositories.BookStoreDBInstance(db)
	services.SetBookstoreInterface(bookstoreInterface)

	routes.BookStoreRoutes(e)
	log.Fatal(e.Start(":9030"))
}
