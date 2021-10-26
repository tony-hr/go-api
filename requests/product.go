package requests

import "github.com/labstack/echo"

type CreateProductRequest struct {
	Name   string `json:"name" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Stock  int    `json:"stock" validate:"required"`
	Images string `json:"images"`
}

func (r *CreateProductRequest) BindValidate(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type EditProductRequest struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Type   string `json:"type" validate:"required"`
	Price  int    `json:"price" validate:"required"`
	Stock  int    `json:"stock" validate:"required"`
	Images string `json:"images"`
}

func (r *EditProductRequest) BindValidate(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type DeleteProductRequest struct {
	ID int `json:"id" validate:"required"`
}

func (r *DeleteProductRequest) BindValidate(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}
