package file

import (
	"net/http"
	"strconv"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateFile(c *gin.Context) {
	var fileDetails models.CreateFileRequest

	if err := c.ShouldBindJSON(&fileDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file := models.File{
		UserID:   uint(fileDetails.UserID),
		Name:     fileDetails.FileName,
		FileSize: uint(fileDetails.FileSize),
		FileType: fileDetails.FileType,
		FolderID: fileDetails.FolderID,
	}

	err := h.service.CreateFile(file)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created file"})
}

func (h *Handler) GetFileByUserID(c *gin.Context) {
	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, serr := h.service.GetFilesByUserID(uint(userId))

	if serr != nil {
		c.JSON(serr.StatusCode, gin.H{"error": serr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (h *Handler) GetFileByUserIDFolderID(c *gin.Context) {

	userIdStr := c.Param("userId")
	folderIdStr := c.Param("folderId")

	userId, err := strconv.Atoi(userIdStr)
	folderId, err := strconv.Atoi(folderIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, serr := h.service.GetFilesByUserIDFolderID(uint(userId), uint(folderId))

	if serr != nil {
		c.JSON(serr.StatusCode, gin.H{"error": serr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
