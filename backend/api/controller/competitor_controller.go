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

	g.GET("", cc.getCompetitors)
	g.POST("/image", cc.saveImage)
}

func (CompetitorController) getCompetitors(c echo.Context) error {
	competitorId := helper.ParseInt(c.QueryParam("id"), 0)
	matchId := helper.ParseInt(c.QueryParam("matchId"), 0)

	match, err := service.MatchService.FindOne(matchId)
	if competitorId == 0 || matchId == 0 || err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	criteria := new(model.Competitor)
	criteria.ID = competitorId
	criteria.MatchID = match.ID

	competitors := make([]model.Competitor, 16)
	if _, err := service.CompetitorService.FindLatest(&competitors, criteria); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, competitors)
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
	matchId := uint64(values["matchId"].(float64))
	match, err := service.MatchService.FindOne(matchId)
	if matchId == 0 || err != nil {
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
