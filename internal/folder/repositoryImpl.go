package folder

import (
	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) CreateFolder(folder models.Folder) error {
	result := r.db.Create(&folder)
	return result.Error
}

func (r *repositoryImpl) GetFoldersByUserID(userId uint) ([]models.Folder, error) {
	var folders []models.Folder

	result := r.db.Where("user_id = ?", userId).Find(&folders)

	if result.Error != nil {
		return nil, result.Error
	}

	return folders, nil
}
