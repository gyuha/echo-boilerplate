package jsonHelper

import (
	"net/http"

	"github.com/labstack/echo"
)

// Message JSON 형태 리턴
func Message(c echo.Context, result bool, message string) error {
	if result == false {
		println(message)
	}
	r := struct {
		Result  bool
		Message string
	}{
		Result:  result,
		Message: message,
	}
	return c.JSON(http.StatusOK, r)
}

// ReturnData 데이터 리턴하기
func Data(c echo.Context, result bool, data interface{}) error {
	r := struct {
		Result bool
		Data   interface{}
	}{
		Result: result,
		Data:   data,
	}
	return c.JSON(http.StatusOK, r)
}
