package user

import (
	"net/http"

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

func (h *Handler) CreateUser(c *gin.Context) {
	var authRequest models.AuthRequest
	var res models.APIResponse

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user := models.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}

	err := h.service.CreateUser(user)

	if err != nil {
		res.Message = err.Message
		res.Success = false
		c.JSON(err.StatusCode, res)
		return
	}

	res.Message = "Successfully created user"
	res.Success = true
	c.JSON(http.StatusCreated, res)
}

func (h *Handler) Login(c *gin.Context) {
	var loginRequest models.AuthRequest
	var res models.APIResponse

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		res.Message = err.Error()
		res.Success = false
		c.JSON(http.StatusBadRequest, res)
		return
	}

	userDto, err := h.service.LoginUser(loginRequest)
	if err != nil {
		res.Message = err.Message
		res.Success = false
		c.JSON(err.StatusCode, res)
		return
	}

	c.SetCookie(
		"token",
		userDto.Token,
		3600,
		"/",
		"localhost",
		false,
		true,
	)

	res.Message = "Successfull logged in"
	res.Success = true
	res.Data = userDto
	c.JSON(http.StatusOK, res)
}
