package file

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Service struct {
	repo   Repository
	client *s3.Client
}

func NewService(repo Repository, client *s3.Client) *Service {
	return &Service{
		repo:   repo,
		client: client,
	}
}

func (s *Service) CreateFile(file multipart.File, fileHeader *multipart.FileHeader, folderId *uint, userId uint) *models.ServiceError {

	// get content type
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	contentType := http.DetectContentType(buf)

	file.Seek(0, io.SeekStart)

	// upload to s3
	_, err = s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:            aws.String(os.Getenv("BUCKET_NAME")),
		Key:               aws.String(fileHeader.Filename),
		ChecksumAlgorithm: types.ChecksumAlgorithmCrc32,
		Body:              file,
		ContentType:       aws.String(contentType),
	})

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// create db entry
	dbFile := models.File{
		UserID:   userId,
		Name:     fileHeader.Filename,
		FileSize: uint(fileHeader.Size),
		FileType: contentType,
		FolderID: folderId,
	}

	err = s.repo.CreateFile(dbFile)

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (s *Service) GetFilesByUserID(userId uint) ([]models.FileDTO, *models.ServiceError) {
	files, err := s.repo.GetFilesByUserID(userId)

	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var filesResponse []models.FileDTO
	for _, file := range files {
		filesResponse = append(filesResponse, models.FileDTO{
			Name:      file.Name,
			Path:      file.Path,
			FileType:  file.FileType,
			FileSize:  file.FileSize,
			UserID:    file.UserID,
			FolderID:  file.FolderID,
			CreatedAt: file.CreatedAt,
		})
	}

	return filesResponse, nil
}

func (s *Service) GetFilesByUserIDFolderID(userId uint, folderId uint) ([]models.FileDTO, *models.ServiceError) {
	files, err := s.repo.GetFilesByUserIDFolderID(userId, folderId)

	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	var filesResponse []models.FileDTO
	for _, file := range files {
		filesResponse = append(filesResponse, models.FileDTO{
			Name:      file.Name,
			Path:      file.Path,
			FileType:  file.FileType,
			FileSize:  file.FileSize,
			UserID:    file.UserID,
			FolderID:  file.FolderID,
			CreatedAt: file.CreatedAt,
		})
	}

	return filesResponse, nil
}
