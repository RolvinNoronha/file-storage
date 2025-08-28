package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Path     string `gorm:"type:varchar(255);not null"`
	FileType string `gorm:"type:varchar(100)"`
	FileSize uint
	UserID   uint   `gorm:"not null"`
	FolderID *uint  `gorm:"index"`
	Folder   Folder `gorm:"foreignKey:FolderID"`
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

type CreateFileRequest struct {
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
	FileType string `json:"fileType"`
	FileSize uint   `json:"fileSize"`
	UserID   uint   `json:"userId"`
	FolderID *uint  `json:"folderId"`
}
