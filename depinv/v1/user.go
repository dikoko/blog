package depinv

import (
	"fmt"

	"github.com/dikoko/blog/depinv/v1/entity"
	"github.com/dikoko/blog/depinv/v1/strg"
)

// UserManager defines the user data manager.
type UserManager struct {
	storage strg.Storage
}

// NewUserManager creates a new UserManager.
func NewUserManager(storage strg.Storage) *UserManager {
	return &UserManager{
		storage: storage,
	}
}

// AddUser adds a user data.
func (m *UserManager) AddUser(user *entity.User) error {
	if user == nil || user.ID == "" {
		return fmt.Errorf("invalid user")
	}
	return m.storage.Add(user.ID, user)
}

// GetUser retrieves a user data.
func (m *UserManager) GetUser(id string) (*entity.User, error) {
	item, err :=  m.storage.Get(id)
	if err != nil {
		return nil, err
	}
	user, ok := item.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("invalid user data")
	}
	return user, nil
}

// DeleteUser deletes a user data.
func (m *UserManager) DeleteUser(id string) error {
	return m.storage.Delete(id)
}
