package controller

import (
	"github.com/labstack/echo"
	"golang-server/helper"
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
	file, err := c.FormFile("zzal")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	contentType := file.Header.Get("Content-Type")
	c.Logger().Info(contentType)
	if !(contentType == "image/jpeg" || contentType == "image/png") {
		return c.NoContent(http.StatusBadRequest)
	}

	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	path, err := service.CompetitorService.SaveFile(src, filepath.Ext(file.Filename))
	strPath, pathErr := path.StringPath()
	if err != nil || pathErr != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	values, err := helper.ConvertJsonToMap(c.FormValue("zzal"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	match, err := service.MatchService.FindOneByMatchName(values["matchName"].(string))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	competitor := &model.Competitor{
		ImageUrl: strPath,
		MatchID:  match.ID,
	}
	if _, err := service.CompetitorService.Save(competitor); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, competitor)
}
