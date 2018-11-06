package controller

import (
	"github.com/labstack/echo"
	"golang-server/helper"
	"golang-server/model"
	"golang-server/service"
	"net/http"
)

type MatchController struct{}

func (mc MatchController) Init(g *echo.Group) {
	g = g.Group("/match")

	g.GET("", mc.getAllMatches)
	g.POST("", mc.createMatch)
	g.GET("/:id", mc.getMatch)
}

func (MatchController) getAllMatches(c echo.Context) error {
	pageable := new(model.Pageable).Of(c.QueryParams())
	matches, err := service.MatchService.FindAll(pageable)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, matches)
}

func (MatchController) createMatch(c echo.Context) error {
	match := new(model.Match)
	if err := c.Bind(match); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if _, err := service.MatchService.Save(match); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (MatchController) getMatch(c echo.Context) error {
	id := helper.ParseInt(c.Param("id"), 0)
	match, err := service.MatchService.FindOne(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, match)
}
