package user

import "github.com/RolvinNoronha/fileupload-backend/pkg/models"


type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(user models.User) (int, error) {
	return s.repo.CreateUser(user);
}
