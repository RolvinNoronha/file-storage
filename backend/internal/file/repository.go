package file

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateFile(models.File) error
	GetFilesByUserID(uint) ([]models.File, error)
	GetFilesByUserIDFolderID(uint, uint) ([]models.File, error)
	GetFile(uint) (*models.File, error)
	UpdateFile(models.File) error
}
