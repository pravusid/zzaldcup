package controller

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"golang-server/model"
	"golang-server/service"
	"net/http"
)

type MatchController struct{}

func (mc MatchController) Init(g *echo.Group) {
	g = g.Group("/match")

	g.GET("", mc.getAllMatches)
	g.POST("", mc.createMatch)
	g.GET("/detail/:matchName", mc.getMatch)
	g.GET("/user", mc.getMatchesOfUser)
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
	if err := c.Validate(match); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if match.Private {
		privateMatch := model.PrivateMatch{
			Match: *match,
			UUID:  uuid.NewV4().String(),
		}
		if err := service.MatchService.SavePrivate(&privateMatch); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

	} else {
		if err := service.MatchService.Save(match); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusCreated)
}

func (MatchController) getMatch(c echo.Context) error {
	matchName := c.Param("matchName")
	var match *model.Match
	var err error
	if c.QueryParam("related") == "true" {
		match, err = service.MatchService.FindOneAndRelatedByMatchName(matchName)
	} else {
		match, err = service.MatchService.FindOneByMatchName(matchName)
	}
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, match)
}

func (MatchController) getMatchesOfUser(c echo.Context) error {
	pageable := new(model.Pageable).Of(c.QueryParams())
	matches, err := service.MatchService.FindUserMatches(pageable)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, matches)
}
