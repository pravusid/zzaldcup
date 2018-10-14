package api

import (
	"github.com/labstack/echo"
	"golang-server/api/controller"
)

func Routes(e *echo.Echo) {
	g := e.Group("/api")

	controller.MatchController{}.Init(g)
	controller.CompetitorController{}.Init(g)
	controller.PlayingController{}.Init(g)
}
