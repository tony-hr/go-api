package repository

import (
	"go-api/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(*models.User) error
	GetByUsername(string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user *models.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
