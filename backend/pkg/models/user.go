package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `gorm:"type:varchar(100);not null"`
	Password string   `gorm:"type:varchar(255)"`
	Files    []File   `gorm:"foreignKey:UserID"`
	Folders  []Folder `gorm:"foreignKey:UserID"`
}

type UserDTO struct {
	Username string `json:"userName"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
}
