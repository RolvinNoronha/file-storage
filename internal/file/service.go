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

func (s *Service) GetFilesByUserID(userId uint) ([]models.File, *models.ServiceError) {
	files, err := s.repo.GetFilesByUserID(userId);

	if (err != nil) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return files, nil;
}

func (s *Service) GetFilesByFolderID(userId uint, folderId uint) ([]models.File, *models.ServiceError) {
	files, err := s.repo.GetFilesByFolderID(userId, folderId);

	if (err != nil) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return files, nil;
}
