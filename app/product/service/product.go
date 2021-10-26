package service

import (
	"errors"
	"go-api/app/product/repository"
	"go-api/helpers"
	"go-api/models"
	"go-api/requests"
	"strings"
)

type IProductService interface {
	AllListProduct() ([]models.Product, error)
	ListProductByID(string) (models.Product, error)
	CreateProduct(*requests.CreateProductRequest) error
	EditProduct(req *requests.EditProductRequest) error
	DeleteProduct(int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) IProductService {
	return &productService{
		repo: repo,
	}
}

func (u *productService) AllListProduct() ([]models.Product, error) {
	allProduct, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return allProduct, nil
}

func (u *productService) ListProductByID(id string) (models.Product, error) {
	product, err := u.repo.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}

	if product == nil {
		return models.Product{}, errors.New("No data available for ID " + id)
	}

	return *product, nil
}

func (u *productService) CreateProduct(req *requests.CreateProductRequest) error {
	if req.Images == "" {
		return errors.New("Image is required")
	}

	imageLower := strings.ToLower(req.Name)
	filename := strings.ReplaceAll(imageLower, " ", "_") + ".png"
	if err := helpers.Base64toImage(req.Images, filename); err != nil {
		return err
	}

	data := models.Product{
		Name:   req.Name,
		Type:   req.Type,
		Price:  req.Price,
		Stock:  req.Stock,
		Images: filename,
	}

	if err := u.repo.Create(&data); err != nil {
		return err
	}

	return nil
}

func (u *productService) EditProduct(req *requests.EditProductRequest) error {
	data := models.Product{
		Name:  req.Name,
		Type:  req.Type,
		Price: req.Price,
		Stock: req.Stock,
	}

	if req.Images != "" {
		imageLower := strings.ToLower(req.Name)
		filename := strings.ReplaceAll(imageLower, " ", "_") + ".png"
		if err := helpers.Base64toImage(req.Images, filename); err != nil {
			return err
		}

		data.Images = filename
	}

	if err := u.repo.Update(req.ID, &data); err != nil {
		return err
	}

	return nil
}

func (u *productService) DeleteProduct(id int) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
