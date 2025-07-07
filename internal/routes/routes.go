package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewRouter(db *gorm.DB) *gin.Engine {
	g := gin.Default()

	return g;
}
