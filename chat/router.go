package chat

import (
	"github.com/labstack/echo"
	melody "gopkg.in/olahol/melody.v1"
)

var (
	m = melody.New()
)

// Router : API에서 사용되는 라우터
func Router(g *echo.Group) {
	// ##############
	// Auth
	g.GET("", ChannelSelect)
	g.GET("/chat/:channel", Chat)

	g.GET("/channel/:channel", HandleRequest)

	m.HandleMessage(HandleMessage)
}
