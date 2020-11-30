package user

import (
	"fmt"

	"github.com/dikoko/blog/depinv/v1/entity"
	"github.com/dikoko/blog/depinv/v1/strg"
)

// Manager defines the user data manager.
type Manager struct {
	storage strg.Storage
}

// NewManager creates a new Manager.
func NewManager(storage strg.Storage) *Manager {
	return &Manager{
		storage: storage,
	}
}

// AddUser adds a user data.
func (m *Manager) AddUser(user *entity.User) error {
	if user == nil || user.ID == "" {
		return fmt.Errorf("invalid user")
	}
	return m.storage.Add(user.ID, user)
}

// GetUser retrieves a user data.
func (m *Manager) GetUser(id string) (*entity.User, error) {
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
func (m *Manager) DeleteUser(id string) error {
	return m.storage.Delete(id)
}
