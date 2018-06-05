package chat

import (
	"github.com/labstack/echo"
)

// Router : API에서 사용되는 라우터
func Router(g *echo.Group) {
	// ##############
	// Auth
	g.GET("", Chat)
	g.File("/chat", "public/chat.html")
}
