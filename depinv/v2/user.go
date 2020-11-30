package user

import (
	"fmt"

	"github.com/dikoko/blog/depinv/v2/entity"
	"github.com/dikoko/blog/depinv/v2/strg"
)

type userStorage interface {
	add(key string, item interface{}) error
	get(key string) (interface{}, error)
	delete(key string) error
}

type userStorageHandler struct {
	storage strg.Storage
}

func (s *userStorageHandler) add(key string, item interface{}) error {
	return s.storage.Add(key, item, "")
}

func (s *userStorageHandler) get(key string) (interface{}, error) {
	item, _, err := s.storage.Get(key)
	return item, err
}

func (s *userStorageHandler) delete(key string) error {
	_, err := s.storage.Delete(key)
	return err
}

// Manager defines the user data manager.
type Manager struct {
	storage userStorage
}

// NewManager creates a new Manager.
func NewManager(storage strg.Storage) *Manager {
	return &Manager{
		storage: &userStorageHandler{
			storage: storage,
		},
	}
}

// AddUser adds a user data.
func (m *Manager) AddUser(user *entity.User) error {
	if user == nil || user.ID == "" {
		return fmt.Errorf("invalid user")
	}
	return m.storage.add(user.ID, user)
}

// GetUser retrieves a user data.
func (m *Manager) GetUser(id string) (*entity.User, error) {
	item, err :=  m.storage.get(id)
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
	return m.storage.delete(id)
}

