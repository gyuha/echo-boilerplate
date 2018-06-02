package api

import (
	"echo-boilerplate/controllers/api/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JwtConfig : main에서 injection으로 넘겨 준다.
var JwtConfig middleware.JWTConfig

// Router : API에서 사용되는 라우터
func Router(g *echo.Group) {
	// ##############
	// Auth
	g.POST("/auth/signup", auth.SignUp)
	g.GET("/auth/users", auth.Users)
	g.POST("/auth/login", auth.Login)
}
