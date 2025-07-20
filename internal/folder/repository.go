
package folder

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateFolder(models.Folder) (error);
	GetFoldersByUserID(uint) ([]models.Folder, error);
}
