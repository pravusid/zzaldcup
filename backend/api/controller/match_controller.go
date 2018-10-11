package controller

import (
	"github.com/labstack/echo"
	"golang-server/helper"
	"golang-server/model"
	"golang-server/service"
	"net/http"
)

var MatchController = &matchController{}

type matchController struct{}

func (mc matchController) Routes(g *echo.Group) {
	g = g.Group("/match")

	g.GET("", mc.getAllMatches)
	g.POST("", mc.createMatch)
	g.GET("/:id", mc.getMatch)
}

func (mc matchController) getAllMatches(c echo.Context) error {
	pageable := new(model.Pageable).Of(c.QueryParams())
	matches, err := service.MatchService.FindAll(pageable)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, matches)
}

func (mc matchController) createMatch(c echo.Context) error {
	match := new(model.Match)
	if err := c.Bind(match); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	service.MatchService.Save(match)
	return c.NoContent(http.StatusCreated)
}

func (mc matchController) getMatch(c echo.Context) error {
	id := helper.ParseInt(c.Param("id"), 0)
	match, err := service.MatchService.FindOne(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, match)
}
