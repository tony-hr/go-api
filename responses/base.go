package responses

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BaseSuccess(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func BaseSuccessWithoutData(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, ResponseWithoutData{
		Code:    http.StatusOK,
		Message: message,
	})
}

func BaseUnauthorized(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, ResponseWithoutData{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}
