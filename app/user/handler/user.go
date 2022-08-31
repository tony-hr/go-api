package handler

import (
	"go-api/app/user/service"
	"go-api/requests"
	"go-api/responses"

	"github.com/labstack/echo"
)

type UserHandler struct {
	Service service.IUserService
}

func (u *UserHandler) Login(c echo.Context) error {
	req := new(requests.LoginRequest)

	if err := req.BindValidate(c); err != nil {
		return err
	}

	respLogin, err := u.Service.Login(req)
	if err != nil {
		return responses.BaseUnauthorized(c, err.Error())
	}

	return responses.BaseSuccess(c, "Login successfully", respLogin)
}
