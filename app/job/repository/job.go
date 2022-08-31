package repository

import (
	"go-api/models"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	GetAll() ([]models.Job, error)
	GetJobByID(string) (*models.Job, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &productRepository{db}
}

func (u *productRepository) GetAll() ([]models.Job, error) {
	var product []models.Job

	return product, nil
}

func (u *productRepository) GetJobByID(id string) (*models.Job, error) {
	var product models.Job
	if err := u.db.Where("id = ?", id).First(&product).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
