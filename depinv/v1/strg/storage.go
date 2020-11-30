package strg

import (
	"fmt"
)

// Storage defines a storage interface.
type Storage interface {
	Add(key string, item interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type storage struct {
	items map[string]interface{}
}

// NewStorage provides a Storage interface.
func NewStorage() Storage {
	return &storage{
		items: make(map[string]interface{}),
	}
}

// Add adds a new item.
func (s *storage) Add(key string, item interface{}) error {
	s.items[key] = item
	return nil
}

// Get retrieves an item.
func (s *storage) Get(key string) (interface{}, error) {
	user, ok := s.items[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return user, nil
}

// Delete deletes an item.
func (s *storage) Delete(key string) error {
	_, ok := s.items[key]
	if !ok {
		return fmt.Errorf("not found")
	}
	delete(s.items, key)
	return nil
}
