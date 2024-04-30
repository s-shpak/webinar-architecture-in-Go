package config

import "hexagonal/internal/adapters/store"

type Config struct {
	Store store.Config
}
