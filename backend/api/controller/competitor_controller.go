package controller

import (
	"github.com/labstack/echo"
	"golang-server/helper"
	"golang-server/model"
	"golang-server/service"
	"net/http"
	"path/filepath"
	"strings"
)

type CompetitorController struct{}

func (cc CompetitorController) Init(g *echo.Group) {
	g = g.Group("/competitor")

	g.GET("", cc.getCompetitors)
	g.PUT("/:id", cc.updateCompetitor)
	g.DELETE("/:id", cc.deleteCompetitor)
	g.GET("/image/:shard/:imageName", cc.getImage)
	g.POST("/image", cc.saveImage)
}

func (CompetitorController) getCompetitors(c echo.Context) error {
	competitorId := helper.ParseInt(c.QueryParam("id"), 0)
	matchId := helper.ParseInt(c.QueryParam("matchId"), 0)

	match, err := service.Match.FindOne(matchId)
	if competitorId == 0 || matchId == 0 || err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	criteria := new(model.Competitor)
	criteria.ID = competitorId
	criteria.MatchID = match.ID

	competitors := make([]model.Competitor, 16)
	if _, err := service.Competitor.FindLatest(&competitors, criteria); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, competitors)
}

func (CompetitorController) updateCompetitor(c echo.Context) error {
	competitorId := helper.ParseInt(c.Param("id"), 0)
	updated := new(model.Competitor)
	if err := c.Bind(updated); err != nil || competitorId != updated.ID {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := service.Competitor.UpdateOne(updated); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (CompetitorController) deleteCompetitor(c echo.Context) error {
	competitor := new(model.Competitor)
	competitor.ID = helper.ParseInt(c.Param("id"), 0)
	if competitor.ID == 0 {
		return c.NoContent(http.StatusInternalServerError)
	}

	err := service.Competitor.DeleteOne(competitor)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (CompetitorController) getImage(c echo.Context) error {
	shard := c.Param("shard")
	imageName := c.Param("imageName")
	return c.File(strings.Join([]string{"image", shard, imageName}, string(filepath.Separator)))
}

func (CompetitorController) saveImage(c echo.Context) error {
	file, err := c.FormFile("zzal")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	contentType := file.Header.Get("Content-Type")
	if !(contentType == "image/jpeg" || contentType == "image/png") {
		return c.NoContent(http.StatusBadRequest)
	}

	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	path, err := service.Competitor.SaveFile(src, filepath.Ext(file.Filename))
	strPath, pathErr := path.StringPath()
	if err != nil || pathErr != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	values, err := helper.ConvertJsonToMap(c.FormValue("zzal"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	matchId := uint64(values["matchId"].(float64))
	match, err := service.Match.FindOne(matchId)
	if matchId == 0 || err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	competitor := &model.Competitor{
		ImageUrl: strPath,
		MatchID:  match.ID,
	}
	if err := service.Competitor.SaveOne(competitor, match); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, competitor)
}
