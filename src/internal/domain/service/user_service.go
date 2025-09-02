package service

import (
	"domain-driven-go/src/internal/domain/model"
	"domain-driven-go/src/internal/domain/repository"
)

// UserService provides user-related services.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(name string) (*model.User, error) {
	user := &model.User{Name: name}
	return user, s.repo.Save(user)
}

// GetUser retrieves a user by their ID.
func (s *UserService) GetUser(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}
