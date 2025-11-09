package file

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	ps *gorm.DB
	es *elasticsearch.Client
}

func NewRepository(ps *gorm.DB, es *elasticsearch.Client) Repository {
	return &repositoryImpl{ps: ps, es: es}
}

/*

func (r *repositoryImpl) insertFileElasticsearch(file models.File) error {
	ctx := context.Background()

	data, err := json.Marshal(file)
	if err != nil {
		return fmt.Errorf("could not marshal doc: %w", err)
	}
	fmt.Print("Inserting into elastic search from after create")

	docID := strconv.Itoa(int(file.ID))
	res, err := r.es.Index(
		"files", // The name of your index
		bytes.NewReader(data),
		r.es.Index.WithDocumentID(docID),
		r.es.Index.WithContext(ctx),
		r.es.Index.WithRefresh("true"), // Make it searchable immediately (slower, remove for bulk)
	)

	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch indexing error: %s", res.String())
	}

	return nil
}

func (r *repositoryImpl) updateFileElasticsearch(file models.File) error {
	ctx := context.Background()

	data, err := json.Marshal(file)
	if err != nil {
		return fmt.Errorf("could not marshal doc: %w", err)
	}

	docID := strconv.Itoa(int(file.ID))
	res, err := r.es.Update(
		"files",
		docID,
		bytes.NewReader(data),
		r.es.Update.WithContext(ctx),
		r.es.Update.WithRefresh("true"),
	)

	if err != nil {
		return fmt.Errorf("error updaing document: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch updating error: %s", res.String())
	}

	return nil
}

func (r *repositoryImpl) deleteFileElasticsearch(fileId uint) error {
	ctx := context.Background()

	docID := strconv.Itoa(int(fileId))
	res, err := r.es.Delete(
		"files",
		docID,
		r.es.Delete.WithContext(ctx),
		r.es.Delete.WithRefresh("true"),
	)

	if err != nil {
		return fmt.Errorf("error updaing document: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch updating error: %s", res.String())
	}

	return nil
}

func (r *repositoryImpl) AfterCreate(tx *gorm.DB) (err error) {
	var file models.File
	if err := tx.Scan(&file).Error; err != nil {
		return err
	}

	fmt.Print("Inserting into elastic search from after create")
	return r.insertFileElasticsearch(file)
}

func (r *repositoryImpl) AfterUpdate(tx *gorm.DB) (err error) {
	var file models.File
	if err := tx.Scan(&file).Error; err != nil {
		return err
	}

	return r.updateFileElasticsearch(file)
}

func (r *repositoryImpl) AfterDelete(tx *gorm.DB) (err error) {
	var file models.File
	if err := tx.Scan(&file).Error; err != nil {
		return err
	}

	return r.deleteFileElasticsearch(file.ID)
}

*/

func (r *repositoryImpl) insertFileElasticsearch(ctx context.Context, fileID uint) error {

	// 1. Fetch the complete data from Postgres
	var file models.File
	// Use Preload to get related Folder data in the same query
	if err := r.ps.WithContext(ctx).Preload("Folder").First(&file, fileID).Error; err != nil {
		// This is a real error; the file *just* got created but we can't find it?
		return fmt.Errorf("could not fetch file for indexing: %w", err)
	}

	// (Optional, but good) Fetch the user to get the username
	var user models.User
	if err := r.ps.WithContext(ctx).First(&user, file.UserID).Error; err != nil {
		log.Printf("Warning: Could not find user %d for file %d: %v", file.UserID, file.ID, err)
		// Don't return an error, just index without the username
	}

	// 2. Create the denormalized document for Elasticsearch
	doc := models.FileDocument{
		ID:         file.ID,
		Name:       file.Name,
		Path:       file.Path,
		FileType:   file.FileType,
		FileSize:   file.FileSize,
		FileUrl:    file.FileUrl,
		CreatedAt:  file.CreatedAt,
		UserID:     file.UserID,
		Username:   user.Username, // <-- Added denormalized data
		FolderID:   file.FolderID,
		FolderName: file.Folder.Name, // <-- Added denormalized data
	}

	// 3. Marshal the 'FileDocument', NOT the 'models.File'
	data, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("could not marshal ES document: %w", err)
	}

	docID := strconv.Itoa(int(doc.ID))
	log.Printf("Indexing document ID: %s", docID) // Use log, not fmt.Print

	// 4. Index in Elasticsearch
	res, err := r.es.Index(
		"files", // The name of your index
		bytes.NewReader(data),
		r.es.Index.WithDocumentID(docID),
		r.es.Index.WithContext(ctx),
		// Use "false" for background indexing. It's much faster.
		// The data will appear "eventually" (usually < 1s).
		r.es.Index.WithRefresh("false"),
	)

	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("elasticsearch indexing error: %s", res.String())
	}

	log.Printf("Successfully indexed file %d", file.ID)
	return nil
}

func (r *repositoryImpl) CreateFile(file models.File) error {
	result := r.ps.Create(&file)
	if result.Error != nil {
		return result.Error
	}
	go func() {
		// We pass the new file's ID.
		// The indexing function will be responsible for fetching all data.
		if err := r.insertFileElasticsearch(context.Background(), file.ID); err != nil {
			// CRITICAL: You must log this error!
			// If this fails, your DB and ES will be out of sync.
			log.Printf("ALERT: Failed to index file %d in Elasticsearch: %v", file.ID, err)
		}
	}()
	return nil
}

func (r *repositoryImpl) GetFilesByUserID(userId uint) ([]models.File, error) {
	var files []models.File

	result := r.ps.Where("user_id = ?", userId).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}

	return files, nil
}

func (r *repositoryImpl) GetFilesByUserIDFolderID(userId uint, folderId uint) ([]models.File, error) {
	var files []models.File

	result := r.ps.Where("user_id = ? AND folder_id = ?", userId, folderId).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}

	return files, nil
}

func (r *repositoryImpl) GetFile(fileId uint) (*models.File, error) {
	var file models.File

	result := r.ps.Where("id = ?", fileId).Find(&file)
	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}

func (r *repositoryImpl) UpdateFile(file models.File) error {
	err := r.ps.Model(&models.File{}).Where("id = ?", file.ID).Updates(models.File{
		FileUrl:       file.FileUrl,
		FileUrlExpiry: file.FileUrlExpiry,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) Search(queryJSON []byte) (*esapi.Response, error) {

	res, err := r.es.Search(
		r.es.Search.WithContext(context.Background()),
		r.es.Search.WithIndex("files"),
		r.es.Search.WithBody(bytes.NewReader(queryJSON)),
		r.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return res, nil
}
