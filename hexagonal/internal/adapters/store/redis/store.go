package redis

import (
	"fmt"
	"hexagonal/internal/core/domain"
)

type Store struct{}

func NewStore(cfg *Config) *Store {
	return &Store{}
}

func (s *Store) GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error) {
	return nil, fmt.Errorf("NYI")
}

func (s *Store) SetFoobar(req *domain.FoobarRequest, resp *domain.FoobarResponse) error {
	return fmt.Errorf("NYI")
}
