package file

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


type Service struct {
	repo Repository
	client *s3.Client
}

func NewService(repo Repository, client *s3.Client) *Service {
	return &Service{
		repo: repo,
		client: client,
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

func (s *Service) GetFilesByUserID(userId uint) ([]models.FileDTO, *models.ServiceError) {
	files, err := s.repo.GetFilesByUserID(userId);

	if (err != nil) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	var filesResponse[] models.FileDTO;
	for _, file := range files {
		filesResponse = append(filesResponse, models.FileDTO{
			Name: file.Name,
			Path: file.Path,
			FileType: file.FileType, 
			FileSize: file.FileSize,  
			UserID: file.UserID,
			FolderID: file.FolderID,
			CreatedAt: file.CreatedAt,
		})
	}

	return  filesResponse, nil;
}

func (s *Service) GetFilesByUserIDFolderID(userId uint, folderId uint) ([]models.FileDTO, *models.ServiceError) {
	files, err := s.repo.GetFilesByUserIDFolderID(userId, folderId);

	if (err != nil) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	var filesResponse[] models.FileDTO;
	for _, file := range files {
		filesResponse = append(filesResponse, models.FileDTO{
			Name: file.Name,
			Path: file.Path,
			FileType: file.FileType, 
			FileSize: file.FileSize,  
			UserID: file.UserID,
			FolderID: file.FolderID,
			CreatedAt: file.CreatedAt,
		})
	}

	return  filesResponse, nil;
}
