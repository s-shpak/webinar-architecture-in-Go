package foobar

import (
	"fmt"
	"strconv"

	"hexagonal/internal/core/domain"
)

type Store interface {
	GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error)
	SetFoobar(req *domain.FoobarRequest, resp *domain.FoobarResponse) error
}

type FoobarSimple struct {
	store Store
}

func NewFoobar(store Store) *FoobarSimple {
	return &FoobarSimple{
		store: store,
	}
}

func (f *FoobarSimple) GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error) {
	resp, err := f.store.GetFoobar(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the foobar result from the store: %w", err)
	}
	if resp != nil {
		return resp, nil
	}
	resp = f.calculateFoobar(req)
	if err := f.store.SetFoobar(req, resp); err != nil {
		return nil, fmt.Errorf("failed to store the foobar calculation in the store: %w", err)
	}
	return resp, nil
}

func (f *FoobarSimple) DoSmth() {
	// do smth
}

func (f *FoobarSimple) calculateFoobar(req *domain.FoobarRequest) *domain.FoobarResponse {
	resp := &domain.FoobarResponse{
		Data: make([]string, 0, req.N),
	}
	for i := 1; i <= req.N; i++ {
		var res string
		if i%3 == 0 && i%5 == 0 {
			res = "foobar"
		} else if i%3 == 0 {
			res = "foo"
		} else if i%5 == 0 {
			res = "bar"
		} else {
			res = strconv.Itoa(i)
		}
		resp.Data = append(resp.Data, res)
	}
	return resp
}
