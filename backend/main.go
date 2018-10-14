package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang-server/api"
	"golang-server/database/mysql"
	"os"
)

func main() {
	defer closeAll()

	e := echo.New()
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.Routes(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func closeAll() {
	mysql.Close()
}
