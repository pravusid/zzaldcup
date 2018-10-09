package controller

import (
	"github.com/labstack/echo"
	"golang-server/model"
	"golang-server/service"
	"net/http"
)

func CompetitorController(g *echo.Group) {
	g = g.Group("/competitor")

	g.GET("", getCompetitors)
	g.POST("", createCompetitors)
	g.POST("/image", saveImage)
}

func getCompetitors(c echo.Context) error {
	return c.String(http.StatusOK, "competitor")
}

func createCompetitors(c echo.Context) error {
	competitors := make([]model.Competitor, 32)
	if err := c.Bind(&competitors); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	service.CompetitorService.Save(&competitors)
	return c.NoContent(http.StatusCreated)
}

func saveImage(c echo.Context) error {
	// TODO: impl
	return nil
}
