package user

import (
	"encoding/json"
	"net/http"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
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
	var authRequest models.AuthRequest;

	reqerr := json.NewDecoder(r.Body).Decode(&authRequest);
	if (reqerr != nil) {
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(map[string]string{"message": "missing required fields!"})
		return;
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)

	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(map[string]string{"message": "could not hash password!"})
		return;
	}

	user := models.User{
		Username: authRequest.Username,
		Password: string(hashedPassword),
	}

	dberr := h.service.repo.GetUserByUsername(user.Username);

	if (dberr != nil) {
		w.WriteHeader(http.StatusConflict);
		json.NewEncoder(w).Encode(map[string]string{"message": "username already exists!"})
		return;
	}


	userId, err := h.service.repo.CreateUser(user);

	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError);
		json.NewEncoder(w).Encode(map[string]string{"message": "failed to add user!"})
		return;
	}

	w.WriteHeader(http.StatusCreated);
	json.NewEncoder(w).Encode(map[string]uint{"userId": userId});
}

 
func (h *Handler) Protected(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK);
}


func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK);
}
