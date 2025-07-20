
package file

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateFile(models.File) (error);
	GetFilesByUserID(uint) ([]models.File, error);
	GetFilesByFolderID(uint, uint) ([]models.File, error);
}
