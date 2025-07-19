package user

import (
	"net/http"
	"time"

	"github.com/RolvinNoronha/fileupload-backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
	jwtSecret []byte
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(user models.User) (*models.ServiceError) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if (err != nil) {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	user.Password = string(hashedPassword);


	_, err = s.repo.GetUserByUsername(user.Username);

	if (err == nil) {
		return &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message: "Username already exists!",
		}
	}


	err = s.repo.CreateUser(user);

	if (err != nil) {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil;
}

func (s *Service) LoginUser(loginRequest models.AuthRequest) (string, *models.ServiceError) {
	
	user, err := s.repo.GetUserByUsername(loginRequest.Username);
	var tokenString string;

	if (err != nil) {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusNotFound,
			Message: "username does not exist.",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password));
	if (err != nil) {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message: "Invalid username or password",
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"expr": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err = token.SignedString([]byte(s.jwtSecret));

	if (err != nil) {
		return tokenString, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}


	return tokenString, nil;
}



