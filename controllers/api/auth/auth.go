package auth

import (
	"errors"
	// "net/http"
	"unicode"

	vali "github.com/asaskevich/govalidator"

	"echo-boilerplate/database/orm"
	"echo-boilerplate/helpers/jsonHelper"
	"echo-boilerplate/models"

	"github.com/labstack/echo"
)

// Users 사용자 목록 부르기
func Users(c echo.Context) error {
	users := []models.User{}
	orm.FindAll(&users)
	return jsonHelper.Data(c, true, users)
}

// verifyPassword : 패스워드 검증
func verifyPassword(s string, min int) bool {
	var (
		number = false
		lower  = false
		// upper   = false
		special = false
		letters = 0
	)
	for _, s := range s {
		switch {
		case unicode.IsLower(s):
			lower = true
		case unicode.IsNumber(s):
			number = true
		// case unicode.IsUpper(s):
		// 	upper = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			special = true
		default:
			//return false, false, false, false
		}
		letters++
	}

	if letters >= min && number == true && lower == true && special == true {
		return true
	}
	return false
}

// signValidate : 로그인용 폼 검증
func signValidate(email string, name string, password string) error {
	if !vali.IsEmail(email) {
		return errors.New("올바른 이메일 형식이 아닙니다")
	}

	if len(name) < 2 {
		return errors.New("별명은 최소 2글자 이상이 필요 합니다")
	}

	if !verifyPassword(password, 6) {
		return errors.New("패스워드는 최소 6자 이상, 영문,숫자,특수 문자를 혼용해야 합니다")
	}

	return nil
}
