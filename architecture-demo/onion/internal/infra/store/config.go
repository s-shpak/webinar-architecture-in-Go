package store

import (
	"onion/internal/infra/store/memory"
	"onion/internal/infra/store/redis"
)

type Config struct {
	Memory *memory.Config
	Redis  *redis.Config
}
