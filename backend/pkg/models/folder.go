package models

import (
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100);not null"`
	UserID         uint   `gorm:"not null"`
	User           User   `gorm:"foreignKey:UserID"`
	ParentFolderID *uint  `gorm:"index;column:parent_folder_id"`
	Files          []File `gorm:"foreignKey:FolderID"`
}

type FolderDTO struct {
	Name      string    `json:"name"`
	UserID    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateFolderRequest struct {
	FolderName     string `json:"folderName"`
	ParentFolderID *uint  `json:"folderId"`
}
