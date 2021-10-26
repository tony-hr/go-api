package requests

import "github.com/labstack/echo"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) BindValidate(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}
