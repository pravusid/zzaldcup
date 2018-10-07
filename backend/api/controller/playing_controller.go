package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PlayingController(g *echo.Group) {
	g = g.Group("/play")

	g.GET("", defaultPlay)
}

func defaultPlay(c echo.Context) error {
	return c.String(http.StatusOK, "Well played!")
}
