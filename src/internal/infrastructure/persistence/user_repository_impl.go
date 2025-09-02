package persistence

import (
	"errors"
	"sync"

	"domain-driven-go/src/internal/domain/model"
	"domain-driven-go/src/internal/domain/repository"
)

// InMemoryUserRepository is an in-memory implementation of the UserRepository.
type InMemoryUserRepository struct {
	mtx    sync.RWMutex
	users  map[int64]*model.User
	nextID int64
}

// NewInMemoryUserRepository creates a new InMemoryUserRepository.
func NewInMemoryUserRepository() repository.UserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int64]*model.User),
		nextID: 1,
	}
}

// Save saves a user to the repository.
func (r *InMemoryUserRepository) Save(user *model.User) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.users[user.ID] = user
	return nil
}

// FindByID finds a user by their ID.
func (r *InMemoryUserRepository) FindByID(id int64) (*model.User, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
