package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang-server/api"
	"golang-server/database"
)

func main() {
	defer database.CloseAll()

	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func customHTTPErrorHandler(err error, e echo.Context) {
	e.Logger().Debug(err)
}
