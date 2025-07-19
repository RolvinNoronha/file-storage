package folder

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
)


type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateFolder(folder models.Folder) (*models.ServiceError) {
	err := s.repo.CreateFolder(folder);

	if (err != nil) {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil;
}
