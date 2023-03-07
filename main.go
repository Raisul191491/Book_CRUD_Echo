package main

import (
	containers "go-bootcamp/pkg/cmd"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Initialize(e)
}
