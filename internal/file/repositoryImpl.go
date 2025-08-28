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

func (r *repositoryImpl) GetFilesByUserID(userId uint) ([]models.File, error) {
	var files[] models.File;

	result := r.db.Where("user_id = ?", userId).Find(&files);
	if (result.Error != nil) {
		return nil, result.Error;
	}

	return files, nil;
}

func (r *repositoryImpl) GetFilesByUserIDFolderID(userId uint, folderId uint) ([]models.File, error) {
	var files[] models.File;

	result := r.db.Where("user_id = ? AND folder_id = ?", userId, folderId).Find(&files);
	if (result.Error != nil) {
		return nil, result.Error;
	}

	return files, nil;
}
