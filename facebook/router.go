package facebook

import (
	"net/http"
	"github.com/labstack/echo"
)

const FACEBOOK_GOPOSTGRES101 = "1701475856573729"

func ShowPage(c echo.Context) error {
	data := struct {
		AppID string
	}{
		AppID: FACEBOOK_GOPOSTGRES101,
	}

	return c.Render(http.StatusOK, "facebookLogin.html", data)
}
// Router : Facebook에서 사용되는 라우터
func Router(g *echo.Group) {
	g.GET("", ShowPage)
}
