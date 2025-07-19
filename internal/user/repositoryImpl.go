package user

import (
	"errors"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db};
}

func (r *repositoryImpl) CreateUser(user models.User) (error) {
	result := r.db.Create(user);

	return result.Error;
}

func (r *repositoryImpl) GetUserByUsername(username string) (models.User, error) {
	var user models.User;
	tx := r.db.Where("username = ?", username)
	result := tx.First(&user);

	if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
		return user, result.Error;
	}

	return user, nil;
}
