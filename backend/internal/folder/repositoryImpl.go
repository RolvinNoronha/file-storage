package folder

import (
	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	ps *gorm.DB
}

func NewRepository(ps *gorm.DB) Repository {
	return &repositoryImpl{ps: ps}
}

func (r *repositoryImpl) CreateFolder(folder models.Folder) error {
	result := r.ps.Create(&folder)
	return result.Error
}

func (r *repositoryImpl) GetFoldersByUserID(userId uint) ([]models.Folder, error) {
	var folders []models.Folder

	result := r.ps.Where("user_id = ?", userId).Find(&folders)

	if result.Error != nil {
		return nil, result.Error
	}

	return folders, nil
}

func (r *repositoryImpl) GetFoldersByFolderID(folderId uint) ([]models.Folder, error) {
	var folders []models.Folder

	result := r.ps.Where("parent_folder_id = ?", folderId).Find(&folders)

	if result.Error != nil {
		return nil, result.Error
	}

	return folders, nil
}
