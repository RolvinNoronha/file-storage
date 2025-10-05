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
	var res models.APIResponse

	userId, exists := c.Get("userId")
	if !exists {
		res.Message = "Request resource is not authorized"
		res.Success = false
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	uidFloat, ok := userId.(float64)

	if !ok {
		res.Message = "Invalide user id type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	var folderDetails models.CreateFolderRequest

	if err := c.ShouldBindJSON(&folderDetails); err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	folder := models.Folder{
		UserID:         uint(uidFloat),
		Name:           folderDetails.FolderName,
		ParentFolderID: folderDetails.ParentFolderID,
	}

	er := h.service.CreateFolder(folder)
	if er != nil {
		res.Message = er.Message
		res.Success = false
		c.JSON(er.StatusCode, res)
		return
	}

	res.Message = "Successfully created folder"
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFolders(c *gin.Context) {
	var res models.APIResponse
	userId, exists := c.Get("userId")

	if !exists {
		res.Message = "Requested resource is not authorized"
		res.Success = false
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	uidFloat, ok := userId.(float64)
	if !ok {
		res.Message = "Invalid user id type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	folders, serr := h.service.GetFolderByUserID(uint(uidFloat))

	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully fetch folders"
	res.Success = true
	res.Data = folders
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFoldersByFolderId(c *gin.Context) {
	var res models.APIResponse
	folderIdStr := c.Param("folderId")

	folderId, err := strconv.Atoi(folderIdStr)
	if err != nil {
		res.Message = "Error parsing folder id"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	folders, serr := h.service.GetFolderByFolderID(uint(folderId))
	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully fetched folders"
	res.Success = true
	res.Data = folders
	c.JSON(http.StatusOK, res)
}
