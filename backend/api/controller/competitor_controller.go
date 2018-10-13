package controller

import (
	"github.com/labstack/echo"
	"golang-server/model"
	"golang-server/service"
	"net/http"
	"path/filepath"
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
	// TODO: content-type chk
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return err
	}

	path, err := service.CompetitorService.SaveFile(src, filepath.Ext(file.Filename))
	if path == "" || err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusCreated, path)
}
