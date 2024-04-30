package application

import (
	"fmt"
	"onion/internal/core/model"
	"onion/internal/core/services"
)

type Repo interface {
	GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error)
}

type Application struct {
	repo Repo
}

func NewApplication(store services.Store) *Application {
	return &Application{
		repo: services.NewRepo(store),
	}
}

func (a *Application) GetFoobar(req *model.FoobarRequest) (*model.FoobarResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("passed foobar request is nil")
	}
	if req.N <= 0 {
		return nil, fmt.Errorf("expected parameter N of the foobar request to be > 0, got %d", req.N)
	}
	return a.repo.GetFoobar(req)
}
