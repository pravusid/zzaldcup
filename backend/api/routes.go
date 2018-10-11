package api

import (
	"github.com/labstack/echo"
	"golang-server/api/controller"
)

func Routes(e *echo.Echo) {
	g := e.Group("/api")

	controller.MatchController.Routes(g)
	controller.CompetitorController.Routes(g)
	controller.PlayingController.Routes(g)
}
