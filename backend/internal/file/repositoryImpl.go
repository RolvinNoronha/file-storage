package file

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

func (r *repositoryImpl) CreateFile(file models.File) error {
	result := r.ps.Create(&file)
	return result.Error
}

func (r *repositoryImpl) GetFilesByUserID(userId uint) ([]models.File, error) {
	var files []models.File

	result := r.ps.Where("user_id = ?", userId).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}

	return files, nil
}

func (r *repositoryImpl) GetFilesByUserIDFolderID(userId uint, folderId uint) ([]models.File, error) {
	var files []models.File

	result := r.ps.Where("user_id = ? AND folder_id = ?", userId, folderId).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}

	return files, nil
}

func (r *repositoryImpl) GetFile(fileId uint) (*models.File, error) {
	var file models.File

	result := r.ps.Where("id = ?", fileId).Find(&file)
	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}

func (r *repositoryImpl) UpdateFile(file models.File) error {
	err := r.ps.Model(&models.File{}).Where("id = ?", file.ID).Updates(models.File{
		FileUrl:       file.FileUrl,
		FileUrlExpiry: file.FileUrlExpiry,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
