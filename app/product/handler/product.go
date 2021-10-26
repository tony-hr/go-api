package handler

import (
	"go-api/app/product/service"
	"go-api/requests"
	"go-api/responses"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	Service service.IProductService
}

func (u *ProductHandler) ListProduct(c echo.Context) error {
	respLogin, err := u.Service.AllListProduct()
	if err != nil {
		return err
	}

	return responses.BaseSuccess(c, "List product has been loaded", respLogin)
}

func (u *ProductHandler) ListProductByID(c echo.Context) error {
	respLogin, err := u.Service.ListProductByID(c.Param("id"))
	if err != nil {
		return err
	}

	return responses.BaseSuccess(c, "Product ID "+c.Param("id"), respLogin)
}

func (u *ProductHandler) CreateProduct(c echo.Context) error {
	req := new(requests.CreateProductRequest)

	if err := req.BindValidate(c); err != nil {
		return err
	}

	if err := u.Service.CreateProduct(req); err != nil {
		return err
	}

	return responses.BaseSuccessWithoutData(c, "Success create product")
}

func (u *ProductHandler) EditProduct(c echo.Context) error {
	req := new(requests.EditProductRequest)

	if err := req.BindValidate(c); err != nil {
		return err
	}

	if err := u.Service.EditProduct(req); err != nil {
		return err
	}

	return responses.BaseSuccessWithoutData(c, "Success update product")
}

func (u *ProductHandler) DeleteProduct(c echo.Context) error {
	req := new(requests.DeleteProductRequest)

	if err := req.BindValidate(c); err != nil {
		return err
	}

	if err := u.Service.DeleteProduct(req.ID); err != nil {
		return err
	}

	return responses.BaseSuccessWithoutData(c, "Delete product successfully")
}
