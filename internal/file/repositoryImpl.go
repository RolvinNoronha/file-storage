package file

import (
	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db};
}

func (r *repositoryImpl) CreateFile(file models.File) (error) {
	result := r.db.Create(&file);
	return result.Error;
}

