package folder

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

func (r *repositoryImpl) CreateFolder(folder models.Folder) (error) {
	result := r.db.Create(&folder);
	return result.Error;
}


