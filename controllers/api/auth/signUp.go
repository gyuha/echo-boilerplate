package auth

import (
	"echo-boilerplate/database/orm"
	"echo-boilerplate/helpers/jsonHelper"
	"echo-boilerplate/models"
	"strings"

	"github.com/labstack/echo"
)

// SignUp : 회원 가입
func SignUp(c echo.Context) error {
	req := struct {
		Email    string
		Password string
		Name     string
	}{}
	if err := c.Bind(&req); err != nil {
		return jsonHelper.Message(c, false, err.Error())
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	if err := signValidate(req.Email, req.Name, req.Password); err != nil {
		return jsonHelper.Message(c, false, err.Error())
	}

	user := new(models.User)

	db := orm.DB()

	db.Where("name = ?", req.Name).First(&user)
	if user.Name != "" {
		return jsonHelper.Message(c, false, "이미 사용중인 별명입니다")
	}

	db.Where("email = ?", req.Email).First(&user)
	if user.Email != "" {
		return jsonHelper.Message(c, false, "이미 사용중인 이메일 입니다")
	}

	tx := db.Begin()
	newUser := models.User{}
	newUser.Email = req.Email
	newUser.Name = req.Name
	newUser.Password, _ = user.PasswordHash(req.Password)

	if err := db.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return jsonHelper.Message(c, false, err.Error())
	}

	tx.Commit()

	// emailHelper.AuthCodeEmailSend(&newUser, req.Email)

	return jsonHelper.Message(c, true, "회원 가입이 완료 되었습니다.")
}
