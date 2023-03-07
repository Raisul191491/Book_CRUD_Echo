package containers

import (
	"go-bootcamp/pkg/app/http/controllers"
	"go-bootcamp/pkg/app/http/routes"
	"go-bootcamp/pkg/app/repositories"
	"go-bootcamp/pkg/app/services"
	"go-bootcamp/pkg/infra/connection"
	"log"

	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo) {

	connection.Connect()
	db := connection.GetDB()

	bookstoreRepo := repositories.BookStoreDBInstance(db)
	services.SetBookstoreInterface(bookstoreRepo)

	bookstoreService := services.BookStoreServiceInstance(bookstoreRepo)
	controllers.SetBookstoreService(bookstoreService)

	routes.BookStoreRoutes(e)
	log.Fatal(e.Start(":9030"))
}
