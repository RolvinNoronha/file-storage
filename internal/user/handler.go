package user

import (
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
)


type Handler struct {
	service *Service
}

func NewHandler(service *Service) (*Handler) {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User;
	_, err := h.service.CreateUser(user);
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError);
	}
	
	w.WriteHeader(http.StatusCreated);
}

 
