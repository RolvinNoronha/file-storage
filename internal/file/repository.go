
package file

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateFile(models.File) (error);
}
