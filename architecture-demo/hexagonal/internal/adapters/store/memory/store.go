package memory

import (
	"hexagonal/internal/core/domain"
	"sync"
)

type Store struct {
	mux *sync.Mutex
	s   map[int][]string
}

func NewStore(cfg *Config) *Store {
	return &Store{
		mux: &sync.Mutex{},
		s:   make(map[int][]string),
	}
}

func (s *Store) GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	res, ok := s.s[req.N]
	if !ok {
		return nil, nil
	}
	return &domain.FoobarResponse{
		Data: res,
	}, nil
}

func (s *Store) SetFoobar(req *domain.FoobarRequest, resp *domain.FoobarResponse) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.s[req.N] = resp.Data
	return nil
}
