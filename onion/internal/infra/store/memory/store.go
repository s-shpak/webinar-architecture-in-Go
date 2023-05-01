package memory

import (
	"onion/internal/core/model"
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

func (s *Store) GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	res, ok := s.s[req.N]
	if !ok {
		return nil, nil
	}
	return &model.FoobarResponse{
		Data: res,
	}, nil
}

func (s *Store) SetFoobar(req *model.FoobarRequest, resp *model.FoobarResponse) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.s[req.N] = resp.Data
	return nil
}
