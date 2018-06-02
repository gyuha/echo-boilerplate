package middlewares

import (
	"echo-boilerplate/helpers/authHelper"
	"net/http"

	"github.com/labstack/echo"
)

// RequireAuth 인증이 필요한 부분
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("AccessToken")
		auth, err := authHelper.ParseTokenString(accessToken)

		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "인증 정보가 없습니다. 다시 인증해 주세요.",
				Internal: err,
			}
		}
		c.Set("User", auth)
		return next(c)
	}
}
