package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func MatchController(g *echo.Group) {
	g = g.Group("/match")

	g.GET("", getMatch)
}

func getMatch(c echo.Context) error {
	return c.String(http.StatusOK, "match")
}
