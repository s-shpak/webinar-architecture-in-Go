package redis

import (
	"fmt"
	"onion/internal/core/model"
)

type Store struct{}

func NewStore(cfg *Config) *Store {
	return &Store{}
}

func (s *Store) GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error) {
	return nil, fmt.Errorf("NYI")
}

func (s *Store) SetFoobar(req *model.FoobarRequest, resp *model.FoobarResponse) error {
	return fmt.Errorf("NYI")
}
