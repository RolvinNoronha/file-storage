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

	var res models.APIResponse

	// check userId
	userId, exists := c.Get("userId")
	if !exists {
		res.Message = "unauthorized request"
		res.Success = false
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	// max file size 15 mb
	err := c.Request.ParseMultipartForm(15 << 20)
	if err != nil {
		res.Message = "File is too large"
		res.Success = false
		c.JSON(http.StatusRequestEntityTooLarge, res)
		return
	}

	folderIdStr := c.Request.FormValue("folderId")

	var folderId *uint = nil
	folderId64, err := strconv.ParseUint(folderIdStr, 10, 64)
	if err == nil {
		fId := uint(folderId64)
		folderId = &fId
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		res.Message = "No file in request"
		res.Success = false
		res.Errors = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	defer file.Close()

	// convert userId to uint
	userIdFloat, ok := userId.(float64)
	if !ok {
		res.Message = "Invalid userId type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	// file upload service
	serr := h.service.CreateFile(file, fileHeader, folderId, uint(userIdFloat))
	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully created file"
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFileByUserID(c *gin.Context) {
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
		res.Message = "Invalid userId type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	files, serr := h.service.GetFilesByUserID(uint(uidFloat))

	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully fetched files"
	res.Data = files
	res.Success = true
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFileByUserIDFolderID(c *gin.Context) {

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
		res.Message = "Invalid userId type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	folderIdStr := c.Param("folderId")
	folderId, err := strconv.Atoi(folderIdStr)

	if err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	files, serr := h.service.GetFilesByUserIDFolderID(uint(uidFloat), uint(folderId))

	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully fetch files"
	res.Success = true
	res.Data = files

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFileUrl(c *gin.Context) {
	var res models.APIResponse
	fileIdStr := c.Param("fileId")

	fileId, err := strconv.Atoi(fileIdStr)

	if err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fileUrl, serr := h.service.GetFileUrl(uint(fileId))

	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfull fetch file url"
	res.Success = true
	res.Data = fileUrl
	c.JSON(http.StatusOK, res)
}

func (h *Handler) SearchFilesHandler(c *gin.Context) {

	var res models.APIResponse
	searchTerm := c.DefaultQuery("q", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}

	searchResult, files, err := h.service.Search(searchTerm, page, size)

	if err != nil {
		res.Message = err.Message
		res.Success = false
		c.JSON(err.StatusCode, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"total": searchResult.Hits.Total.Value,
		"page":  page,
		"size":  size,
		"data":  files,
	})
}

func (h *Handler) InitiateMultipartUpload(c *gin.Context) {
	var req models.InitiateMultipartUploadRequest
	var res models.APIResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		res.Message = "unauthorized request"
		res.Success = false
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	uidFloat, ok := userId.(float64)
	if !ok {
		res.Message = "Invalid userId type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	response, serr := h.service.InitiateMultipartUpload(req, uint(uidFloat))
	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully initiated upload"
	res.Success = true
	res.Data = response
	c.JSON(http.StatusOK, res)
}

func (h *Handler) CompleteMultipartUpload(c *gin.Context) {
	var req models.CompleteMultipartUploadRequest
	var res models.APIResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		res.Message = "unauthorized request"
		res.Success = false
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	uidFloat, ok := userId.(float64)
	if !ok {
		res.Message = "Invalid userId type"
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	serr := h.service.CompleteMultipartUpload(req, uint(uidFloat))
	if serr != nil {
		res.Message = serr.Message
		res.Success = false
		c.JSON(serr.StatusCode, res)
		return
	}

	res.Message = "Successfully completed upload"
	res.Success = true
	c.JSON(http.StatusOK, res)
}
