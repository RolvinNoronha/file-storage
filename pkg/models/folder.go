package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Files  []File `gorm:"foreignKey:FolderID"`
}
