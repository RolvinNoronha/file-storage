package file

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) (*Handler) {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateFile(c *gin.Context) {
	var fileDetails models.CreateFileRequest;

	if err := c.ShouldBindJSON(&fileDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}


	file := models.File{
		UserID: uint(fileDetails.UserID),
	    Name : fileDetails.FileName,
		FileSize: uint(fileDetails.FileSize),
		FileType: fileDetails.FileType,
		FolderID: fileDetails.FolderID,
	}

	err := h.service.CreateFile(file);
	if (err != nil) {
		c.JSON(err.StatusCode, gin.H{"error": err.Message});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created file"});
}

func (h *Handler) GetFileByUserID(c *gin.Context) {
	var getFilesRequest models.GetFilesByUserIDRequest;

	if err := c.ShouldBindJSON(&getFilesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	files, err := h.service.GetFilesByUserID(uint(getFilesRequest.UserID));

	if (err != nil) {
		c.JSON(err.StatusCode, gin.H{"error": err.Message});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"files": files});
}


func (h *Handler) GetFileByFolderID(c *gin.Context) {
	var getFilesRequest models.GetFilesByFolderIDRequest;

	if err := c.ShouldBindJSON(&getFilesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	files, err := h.service.GetFilesByFolderID(uint(getFilesRequest.UserID), uint(getFilesRequest.FolderID));

	if (err != nil) {
		c.JSON(err.StatusCode, gin.H{"error": err.Message});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"files": files});
}


