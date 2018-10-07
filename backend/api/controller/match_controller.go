package controller

import (
	"github.com/labstack/echo"
	"golang-server/model"
	"golang-server/service"
	"strconv"
)

func MatchController(g *echo.Group) {
	g = g.Group("/match")

	g.GET("", getAllMatches)
	g.POST("", createMatch)
	g.GET("/:id", getMatch)
}

func getAllMatches(c echo.Context) error {
	pageable := model.NewPageable(c.QueryParams())
	matches, err := service.MatchService.FindAll(pageable)
	if err != nil {
		return c.NoContent(404)
	}
	return c.JSON(200, matches)
}

func createMatch(c echo.Context) error {
	match := model.Match{}
	if err := c.Bind(&match); err != nil {
		return c.NoContent(400)
	}
	service.MatchService.Save(&match)
	return c.NoContent(201)
}

func getMatch(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	match, err := service.MatchService.FindOne(id)
	if err != nil {
		return c.NoContent(404)
	}
	return c.JSON(200, match)
}
