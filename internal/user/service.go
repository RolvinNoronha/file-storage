package user

import (
	"net/http"
	"os"
	"time"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      Repository
	jwtSecret []byte
}

func NewService(repo Repository) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (s *Service) CreateUser(user models.User) *models.ServiceError {

	existingUser, err := s.repo.GetUserByUsername(user.Username)

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if existingUser != nil {
		return &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Username already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error hashing password",
		}
	}

	user.Password = string(hashedPassword)

	err = s.repo.CreateUser(user)

	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error creating user in database",
		}
	}

	return nil
}

func (s *Service) LoginUser(loginRequest models.AuthRequest) (string, *models.ServiceError) {

	user, err := s.repo.GetUserByUsername(loginRequest.Username)
	var tokenString string

	if user == nil {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusNotFound,
			Message:    "Username does not exist.",
		}
	}

	if err != nil {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusNotFound,
			Message:    "Something went wrong",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid username or password",
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"expr":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err = token.SignedString([]byte(s.jwtSecret))

	if err != nil {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return tokenString, nil
}
