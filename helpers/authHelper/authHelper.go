package authHelper

import (
	"echo-boilerplate/conf"
	"errors"

	"github.com/labstack/echo"

	"github.com/dgrijalva/jwt-go"
)

// AuthClaims 인증 용
type AuthClaims struct {
	jwt.StandardClaims
	ID       uint
	Name     string
	UserCode string
	Email    string
	Role     int
}

// CreateTokenString : 인증정보 만들어 주기
func CreateTokenString(authClaims AuthClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims)
	return token.SignedString([]byte(conf.Conf.App.JwtSecret))
}

// ParseTokenString : 스트링에서 인증 정보를 파싱한다.
func ParseTokenString(tokenString string) (AuthClaims, error) {
	authClaims := AuthClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &authClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.App.JwtSecret), nil
	})
	return authClaims, err
}

// ReadAuthClaimsByCookie : 인증 정보를 쿠기에서 가져 온다.
func ReadAuthClaimsByCookie(c echo.Context) (AuthClaims, error) {
	authClaims := AuthClaims{}
	tokenString, err := c.Cookie("User")
	if err != nil {
		return authClaims, errors.New("didn't read cookie")
	}
	_, err = jwt.ParseWithClaims(tokenString.Value, &authClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.App.JwtSecret), nil
	})
	return authClaims, err
}

// GetCurrentUser 현재 사용자 정보 리턴
func GetCurrentUser(c echo.Context) AuthClaims {
	return c.Get("User").(AuthClaims)
}
