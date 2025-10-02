package file

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Service struct {
	repo          Repository
	client        *s3.Client
	presignClient *s3.PresignClient
	bucketName    string
}

func NewService(repo Repository, client *s3.Client) *Service {
	return &Service{
		repo:          repo,
		client:        client,
		bucketName:    os.Getenv("BUCKET_NAME"),
		presignClient: s3.NewPresignClient(client),
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
		Bucket:            aws.String(s.bucketName),
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

func (s *Service) GetFileById(fileId uint) (*models.File, *models.ServiceError) {

	file, err := s.repo.GetFile(fileId)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return file, nil
}

func (s *Service) UpdateFile(file *models.File) *models.ServiceError {

	err := s.repo.UpdateFile(*file)
	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return nil
}

func (s *Service) GetFileUrl(fileId uint) (*models.FileUrlDTO, *models.ServiceError) {

	file, err := s.GetFileById(fileId)
	if err != nil {
		return nil, err
	}

	if time.Now().After(file.ExpiresAt) {
		bucketName := s.bucketName
		objectKey := file.Name
		expiresIn := 60 * time.Minute // URL valid for 15 minutes
		expiryTime := time.Now().Add(expiresIn)

		presignedURL, err := s.presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &objectKey,
		}, s3.WithPresignExpires(expiresIn))

		if err != nil {
			return nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}

		file.FileUrl = presignedURL.URL
		file.ExpiresAt = expiryTime

		s.UpdateFile(file)

		fileUrl := &models.FileUrlDTO{
			FileUrl: presignedURL.URL,
			FileId:  fileId,
		}

		return fileUrl, nil
	}

	fileUrl := &models.FileUrlDTO{
		FileUrl: file.FileUrl,
		FileId:  fileId,
	}

	return fileUrl, nil
}
