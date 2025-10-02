package folder

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

func (h *Handler) CreateFolder(c *gin.Context) {
	var folderDetails models.CreateFolderRequest

	if err := c.ShouldBindJSON(&folderDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(folderDetails.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	folder := models.Folder{
		UserID: uint(id),
		Name:   folderDetails.FolderName,
	}

	er := h.service.CreateFolder(folder)
	if er != nil {
		c.JSON(er.StatusCode, gin.H{"error": er.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created folder"})
}

func (h *Handler) GetFolders(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Request resource is not authorized"})
		return
	}

	uidFloat, ok := userId.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userId type"})
		return
	}

	folders, serr := h.service.GetFolderByUserID(uint(uidFloat))

	if serr != nil {
		c.JSON(serr.StatusCode, gin.H{"error": serr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"folders": folders})
}

func (h *Handler) GetFoldersByFolderId(c *gin.Context) {
	folderIdStr := c.Param("folderId")

	folderId, err := strconv.Atoi(folderIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing folderID"})
		return
	}

	folders, serr := h.service.GetFolderByFolderID(uint(folderId))
	if serr != nil {
		c.JSON(serr.StatusCode, gin.H{"error": serr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"folders": folders})
}
