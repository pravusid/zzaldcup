package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func CompetitorController(g *echo.Group) {
	g = g.Group("/competitor")

	g.GET("", getCompetitors)
	g.POST("", createCompetitors)
}

func getCompetitors(c echo.Context) error {
	return c.String(http.StatusOK, "competitor")
}

func createCompetitors(c echo.Context) error {
	return c.NoContent(201)
}
