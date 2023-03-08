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

	bookstoreRepo := repositories.BookStoreDBInstance(db)
	services.SetBookstoreInterface(bookstoreRepo)

	bookstoreService := services.BookStoreServiceInstance(bookstoreRepo)
	controllers.SetBookstoreService(bookstoreService)

	routes.BookStoreRoutes(e)
	log.Fatal(e.Start(":9030"))
}
