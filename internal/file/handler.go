package file

import (
	"net/http"
	"strconv"

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

	// check userId
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unorthorized request"})
		return
	}

	// max file size 15 mb
	err := c.Request.ParseMultipartForm(15 << 20)
	if err != nil {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File is too large"})
		return
	}

	folderIdStr := c.Request.FormValue("folderId")
	var folderId64 uint64

	folderId64, _ = strconv.ParseUint(folderIdStr, 10, 64)
	folderId := uint(folderId64)

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	// convert userId to uint
	userIdFloat, ok := userId.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userId type"})
		return
	}

	// file upload service
	serr := h.service.CreateFile(file, fileHeader, &folderId, uint(userIdFloat))
	if serr != nil {
		c.JSON(serr.StatusCode, gin.H{"error": serr.Message})
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
