package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name          string     `gorm:"type:varchar(100);not null"`
	Path          string     `gorm:"type:varchar(255);not null"`
	FilePath      string     `gorm:"type:varchar(255)"`
	FileType      string     `gorm:"type:varchar(100)"`
	FileUrl       string     `gorm:"type:varchar(255)"`
	FileUrlExpiry *time.Time `gorm:"type:date"`
	FileSize      uint
	UserID        uint   `gorm:"not null"`
	FolderID      *uint  `gorm:"index"`
	Folder        Folder `gorm:"foreignKey:FolderID"`
}

type FileDocument struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	FileType   string    `json:"file_type"`
	FileSize   uint      `json:"file_size"`
	FileUrl    string    `json:"file_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     uint      `json:"user_id"`
	Username   string    `json:"username"` // Denormalized
	FolderID   *uint     `json:"folder_id"`
	FolderName string    `json:"folder_name,omitempty"` // Denormalized
}

type FileDTO struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	FileType  string    `json:"type"`
	FileSize  uint      `json:"size"`
	UserID    uint      `json:"userId"`
	FolderID  *uint     `json:"folderId"`
	CreatedAt time.Time `json:"createdAt"`
}

type FileUrlDTO struct {
	FileUrl string `json:"fileUrl"`
	FileId  uint   `json:"fileId"`
}

type CreateFileRequest struct {
	FolderID *uint `json:"folderId"`
}

type InitiateMultipartUploadRequest struct {
	FileName string `json:"fileName" binding:"required"`
	FileType string `json:"fileType" binding:"required"`
	FileSize int64  `json:"fileSize" binding:"required"`
	FolderID *uint  `json:"folderId"`
}

type InitiateMultipartUploadResponse struct {
	UploadId string          `json:"uploadId"`
	Key      string          `json:"key"`
	Parts    []PresignedPart `json:"parts"`
}

type PresignedPart struct {
	PartNumber int    `json:"partNumber"`
	Url        string `json:"url"`
}

type CompleteMultipartUploadRequest struct {
	UploadId string          `json:"uploadId" binding:"required"`
	Key      string          `json:"key" binding:"required"`
	Parts    []CompletedPart `json:"parts" binding:"required"`
	FileName string          `json:"fileName" binding:"required"`
	FileSize uint            `json:"fileSize" binding:"required"`
	FileType string          `json:"fileType" binding:"required"`
	FolderID *uint           `json:"folderId"`
}

type CompletedPart struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"etag"`
}
