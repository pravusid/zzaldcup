package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

var PlayingController = &playingController{}

type playingController struct{}

func (pc playingController) Routes(g *echo.Group) {
	g = g.Group("/play")

	g.GET("", pc.defaultPlay)
}

func (pc playingController) defaultPlay(c echo.Context) error {
	return c.String(http.StatusOK, "Well played!")
}
