package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang-server/api"
	"golang-server/database/mysql"
	"golang-server/model"
	"os"
)

func main() {
	defer closeAll()
	connectAll()

	e := echo.New()
	if os.Getenv("PROFILE") == "dev" {
		e.Debug = true
	}
	e.Validator = model.Validator

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CLIENT_URI")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	api.Routes(e)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func connectAll() {
	mysql.Init()
}

func closeAll() {
	mysql.Close()
}
