package main

import (
	"todo/backend/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	con := controller.New()
	controller.RegisterHandlers(e, con)
}
