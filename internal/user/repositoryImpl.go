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

func (r *repositoryImpl) CreateUser(user models.User) (uint, error) {
	r.db.Create(&user);
	result := r.db.Create(user);

	return user.ID, result.Error;
}

func (r *repositoryImpl) GetUserByUsername(email string) (error) {
	var user models.User;
	tx := r.db.Where("username = ?", email)
	result := tx.First(&user);

	if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
		return result.Error;
	}

	return nil;
}
