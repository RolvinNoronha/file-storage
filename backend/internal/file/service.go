package file

import (
	"context"
	"encoding/json"
	"fmt"
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

	var path string
	if folderId != nil {
		path = fmt.Sprintf("user-files/%d/%d/%s", userId, *folderId, fileHeader.Filename)
	} else {
		// If no folderId is provided, upload to a default folder
		path = fmt.Sprintf("user-files/%d/%s", userId, fileHeader.Filename)
	}

	// upload to s3
	_, err = s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:            aws.String(s.bucketName),
		Key:               aws.String(path),
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
		UserID:        userId,
		FilePath:      path,
		Name:          fileHeader.Filename,
		FileSize:      uint(fileHeader.Size),
		FileType:      contentType,
		FolderID:      folderId,
		FileUrlExpiry: nil,
		FileUrl:       "",
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

	if file.FileUrlExpiry == nil || time.Now().After(*file.FileUrlExpiry) {
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
		file.FileUrlExpiry = &expiryTime

		serr := s.UpdateFile(file)
		if serr != nil {
			return nil, serr
		}

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

func (s *Service) Search(searchTerm string, page int, size int) (*models.SearchResult, []models.FileDocument, *models.ServiceError) {

	// Calculate 'from' for ES pagination
	from := (page - 1) * size

	// 2. Build the Elasticsearch Query
	var query map[string]interface{}
	if searchTerm == "" {
		// Match all if no search term is provided
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
		}
	} else {
		// Use a 'multi_match' query to search multiple fields
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  searchTerm,
					"fields": []string{"name", "file_type", "username", "folder_name"},
				},
			},
		}
	}

	// 3. Add pagination and sorting
	query["from"] = from
	query["size"] = size
	query["sort"] = []interface{}{
		map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}

	// Marshal the query map to a JSON string
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not build query",
		}
	}

	res, err := s.repo.Search(queryJSON)

	// 5. Parse the response
	var searchResult models.SearchResult
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Could not parse search results",
		}
	}

	// 6. Format and return the paginated response
	// Extract the documents from the 'hits'
	files := make([]models.FileDocument, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {
		files[i] = hit.Source
	}

	return &searchResult, files, nil
}
