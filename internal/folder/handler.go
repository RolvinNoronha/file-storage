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

func NewHandler(service *Service) (*Handler) {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateFolder(c *gin.Context) {
	var folderDetails models.CreateFolderRequest;

	if err := c.ShouldBindJSON(&folderDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	id, err := strconv.Atoi(folderDetails.UserID);
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});
		return;
	}

	folder := models.Folder{
		UserID: uint(id),
		Name: folderDetails.FolderName,
	}

	er := h.service.CreateFolder(folder);
	if (er != nil) {
		c.JSON(er.StatusCode, gin.H{"error": er.Message});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created folder"});
}

func (h *Handler) GetFoldersByUserID(c *gin.Context) {
	var getFoldersRequest *models.GetFolderByUserIDRequest;	

	if err := c.ShouldBindJSON(&getFoldersRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}

	folders, err := h.service.GetFolderByUserID(uint(getFoldersRequest.UserID));

	if (err != nil) {
		c.JSON(err.StatusCode, gin.H{"error": err.Message});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"folders": folders});
}

