package controller

import (
	"github.com/labstack/echo"
	"golang-server/model"
	"golang-server/service"
	"io"
	"net/http"
	"os"
)

var CompetitorController = &competitorController{}

type competitorController struct{}

func (cc competitorController) Routes(g *echo.Group) {
	g = g.Group("/competitor")

	g.GET("", cc.getCompetitors)
	g.POST("", cc.createCompetitors)
	g.POST("/image", cc.saveImage)
}

func (cc competitorController) getCompetitors(c echo.Context) error {
	return c.String(http.StatusOK, "competitor")
}

func (cc competitorController) createCompetitors(c echo.Context) error {
	competitors := make([]model.Competitor, 32)
	if err := c.Bind(&competitors); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	service.CompetitorService.Save(&competitors)
	return c.NoContent(http.StatusCreated)
}

func (cc competitorController) saveImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err = io.Copy(dest, src); err != nil {
		return err
	}

	return c.String(http.StatusCreated, "")
}
