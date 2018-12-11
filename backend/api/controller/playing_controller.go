package controller

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"net/http"
)

type PlayingController struct {
	upgrader websocket.Upgrader
}

func (pc PlayingController) Init(g *echo.Group) {
	g = g.Group("/play")

	g.POST("", pc.createPlay)
	g.GET("/{uuid}", pc.loadPlay)

	pc.upgrader = websocket.Upgrader{}
}

func (pc *PlayingController) createPlay(c echo.Context) error {
	return c.String(http.StatusOK, "Well played!")
}

func (pc *PlayingController) loadPlay(c echo.Context) error {
	ws, err := pc.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, World!"))
		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		c.Logger().Infof("%s\n", msg)
	}
}
