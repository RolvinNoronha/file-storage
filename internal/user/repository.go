package user

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"

type Repository interface {
	CreateUser(models.User) error
	GetUserByUsername(string) (*models.User, error)
}
