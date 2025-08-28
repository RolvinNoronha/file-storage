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

func (s *Service) CreateFolder(folder models.Folder) *models.ServiceError {
	err := s.repo.CreateFolder(folder)

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (s *Service) GetFolderByUserID(userId uint) ([]models.FolderDTO, *models.ServiceError) {
	folders, err := s.repo.GetFoldersByUserID(userId)

	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var folderResponse []models.FolderDTO
	for _, folder := range folders {
		folderResponse = append(folderResponse, models.FolderDTO{
			Name:      folder.Name,
			UserID:    folder.UserID,
			CreatedAt: folder.CreatedAt,
		})
	}

	return folderResponse, nil
}
