package foobar

import (
	"fmt"

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
	resp, err = f.calculateFoobar(req)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate the foobar result for %+v: %w", req, err)
	}
	if err := f.store.SetFoobar(req, resp); err != nil {
		return nil, fmt.Errorf("failed to store the foobar calculation in the store: %w", err)
	}
	return nil, nil
}

func (f *FoobarSimple) calculateFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error) {
	return nil, nil
}
