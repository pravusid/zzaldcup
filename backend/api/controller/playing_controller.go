package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type PlayingController struct{}

func (pc PlayingController) Init(g *echo.Group) {
	g = g.Group("/play")

	g.GET("", pc.defaultPlay)
}

func (PlayingController) defaultPlay(c echo.Context) error {
	return c.String(http.StatusOK, "Well played!")
}
