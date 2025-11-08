package user

import (
	"errors"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	ps *gorm.DB
}

func NewRepository(ps *gorm.DB) Repository {
	return &repositoryImpl{ps: ps}
}

func (r *repositoryImpl) CreateUser(user models.User) error {
	result := r.ps.Create(&user)

	return result.Error
}

func (r *repositoryImpl) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	tx := r.ps.Where("username = ?", username)
	result := tx.First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}
