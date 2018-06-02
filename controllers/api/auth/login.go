package auth

import (
	"echo-boilerplate/database/orm"
	"echo-boilerplate/helpers/authHelper"
	"echo-boilerplate/helpers/jsonHelper"
	"echo-boilerplate/models"
	"errors"
	"net/http"
	"time"

	vali "github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// loginValidate : 로그인용 폼 검증
func loginValidate(email string, password string) error {
	if !vali.IsEmail(email) {
		return errors.New("올바른 이메일 형식이 아닙니다")
	}

	if !verifyPassword(password, 6) {
		return errors.New("패스워드는 최소 6자 이상, 영문,숫자,특수 문자를 혼용해야 합니다")
	}

	return nil
}

// Login : 회원 로그인
func Login(c echo.Context) error {
	req := struct {
		Email    string
		Password string
	}{}

	if err := c.Bind(&req); err != nil {
		return jsonHelper.Message(c, false, err.Error())
	}

	if err := loginValidate(req.Email, req.Password); err != nil {
		return jsonHelper.Message(c, false, err.Error())
	}

	user := new(models.User)
	db := orm.DB()
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return jsonHelper.Message(c, false, "아이디와 암호를 다시 확인해 주세요")
	}

	if !user.ComparePassword(user.Password, req.Password) {
		return jsonHelper.Message(c, false, "패스워드가 잘 못 되었습니다")
	}

	// print.Pretty(user)

	if user.Role == "pending" {
		return c.JSON(http.StatusOK, "이메일 인증이 되지 않았습니다. 이메일 인증 후 다시 시도해 주세요.")
	}

	authClaims := authHelper.AuthClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			// ExpiresAt: time.Now().Unix(),
		},
		user.ID,
		user.Name,
		user.UserCode,
		user.Email,
		user.Role,
	}
	accessToken, _ := authHelper.CreateTokenString(authClaims)
	authClaims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 240).Unix()
	refreshToken, _ := authHelper.CreateTokenString(authClaims)

	// cookie := new(http.Cookie)
	// cookie.Name = "user"
	// cookie.Value = jwtString
	// cookie.Expires = time.Now().Add(time.Hour * 72)
	// c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Result":       true,
		"Data":         user,
		"Email":        user.Email,
		"AccessToken":  accessToken,
		"RefreshToken": refreshToken,
	})
}
