package file

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

func (s *Service) CreateFile(file models.File) (*models.ServiceError) {
	err := s.repo.CreateFile(file);

	if (err != nil) {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil;
}

