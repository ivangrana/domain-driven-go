package repository

import "domain-driven-go/src/internal/domain/model"

// UserRepository defines the interface for user data persistence.
type UserRepository interface {
	Save(user *model.User) error
	FindByID(id int64) (*model.User, error)
}
