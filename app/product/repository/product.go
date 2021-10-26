package repository

import (
	"go-api/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(*models.Product) error
	Update(int, *models.Product) error
	Delete(int) error
	GetProductByID(string) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (u *productRepository) Create(product *models.Product) error {
	return u.db.Create(&product).Error
}

func (u *productRepository) Update(id int, product *models.Product) error {
	return u.db.Model(product).Where("id=?", id).Update(product).Error
}

func (u *productRepository) GetAll() ([]models.Product, error) {
	var product []models.Product
	if err := u.db.Find(&product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

func (u *productRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	if err := u.db.Where("id = ?", id).First(&product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (u *productRepository) Delete(ID int) error {
	model := models.Product{}
	return u.db.Where("id = ?", ID).Delete(model).Error
}
