package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name     string   `gorm:"type:varchar(100);not null"`
	Path     string   `gorm:"type:varchar(255);not null"`
	FileType string   `gorm:"type:varchar(100)"` 
	FileSize uint 
	UserID   uint     `gorm:"not null"`
	FolderID *uint    `gorm:"index"` 
	Folder   Folder   `gorm:"foreignKey:FolderID"` 
}
