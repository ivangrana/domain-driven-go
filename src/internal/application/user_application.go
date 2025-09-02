package application

import (
	"domain-driven-go/src/internal/domain/model"
	"domain-driven-go/src/internal/domain/service"
)

// UserApplication provides a high-level API for user-related operations.
type UserApplication struct {
	userService *service.UserService
}

// NewUserApplication creates a new UserApplication.
func NewUserApplication(userService *service.UserService) *UserApplication {
	return &UserApplication{userService: userService}
}

// CreateUser creates a new user.
func (a *UserApplication) CreateUser(name string) (*model.User, error) {
	return a.userService.CreateUser(name)
}

// GetUser retrieves a user by their ID.
func (a *UserApplication) GetUser(id int64) (*model.User, error) {
	return a.userService.GetUser(id)
}
