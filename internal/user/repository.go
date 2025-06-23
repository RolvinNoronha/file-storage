package user

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateUser(models.User) (int, error);
}
