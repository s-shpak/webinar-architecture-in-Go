package services

import (
	"fmt"

	"onion/internal/core/model"
)

type Store interface {
	GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error)
	SetFoobar(req *model.FoobarRequest, resp *model.FoobarResponse) error
}

type foobar interface {
	Calculate(req *model.FoobarRequest) *model.FoobarResponse
}

type Repo struct {
	store  Store
	foobar foobar
}

func NewRepo(store Store) *Repo {
	return &Repo{
		store:  store,
		foobar: &foobarSimple{},
	}
}

func (r *Repo) GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error) {
	resp, err := r.store.GetFoobar(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the foobar result from the store: %w", err)
	}
	if resp != nil {
		return resp, nil
	}
	resp = r.foobar.Calculate(req)
	if err := r.store.SetFoobar(req, resp); err != nil {
		return nil, fmt.Errorf("failed to store the foobar calculation in the store: %w", err)
	}
	return resp, nil
}
