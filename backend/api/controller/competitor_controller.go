package controller

import (
	"github.com/labstack/echo"
	"golang-server/model"
	"golang-server/service"
	"net/http"
	"path/filepath"
)

type CompetitorController struct{}

func (cc CompetitorController) Init(g *echo.Group) {
	g = g.Group("/competitor")

	g.POST("", cc.createCompetitor)
	g.POST("/image", cc.saveImage)
}

func (CompetitorController) createCompetitor(c echo.Context) error {
	competitor := new(model.Competitor)
	if err := c.Bind(competitor); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err := service.MatchService.FindOne(competitor.ID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	service.CompetitorService.Save(competitor)
	return c.NoContent(http.StatusCreated)
}

func (CompetitorController) saveImage(c echo.Context) error {
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
