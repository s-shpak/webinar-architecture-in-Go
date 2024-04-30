package store

import (
	"hexagonal/internal/adapters/store/memory"
	"hexagonal/internal/adapters/store/redis"
)

type Config struct {
	Memory *memory.Config
	Redis  *redis.Config
}
