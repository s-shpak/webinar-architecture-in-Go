package store

import (
	"hexagonal/internal/adapters/store/memory"
	"hexagonal/internal/adapters/store/redis"
	"hexagonal/internal/config"
	"hexagonal/internal/core/domain"
)

type Store interface {
	GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error)
	SetFoobar(req *domain.FoobarRequest, resp *domain.FoobarResponse) error
}

func NewStore(conf config.Config) (Store, error) {
	if conf.Memory != nil {
		return memory.NewStore(conf.Memory), nil
	}
	if conf.Redis != nil {
		return redis.NewStore(conf.Redis), nil
	}
	return nil, nil
}
