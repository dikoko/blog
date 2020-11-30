package strg

import (
	"fmt"
)

type Storage struct {
	items map[string]interface{}
	notes map[string]string
}

// NewStorage provides a Storage interface.
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]interface{}),
		notes: make(map[string]string),
	}
}

// Add adds a new item.
func (s *Storage) Add(key string, item interface{}, note string) error {
	s.items[key] = item
	s.notes[key] = note
	return nil
}

// Get retrieves an item.
func (s *Storage) Get(key string) (interface{}, string, error) {
	user, ok := s.items[key]
	if !ok {
		return nil, "", fmt.Errorf("not found")
	}
	return user, s.notes[key], nil
}

// Delete deletes an item.
func (s *Storage) Delete(key string) (interface{}, error) {
	item, ok := s.items[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	delete(s.items, key)
	return item, nil
}
