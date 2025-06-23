package user

import (
	"fmt"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db:db};
}

func (r *RepositoryImpl) CreateUser(user models.User) (int, error) {
	userId := r.db.Create(user);

	fmt.Print(userId);

	return 2, nil
}
