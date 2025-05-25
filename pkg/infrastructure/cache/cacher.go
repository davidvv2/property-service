package redis

import (
	"context"
	"time"
)

// Cacher defines the interface for Redis caching operations.
type Cacher interface {
	KeySet(ctx context.Context, key string, value interface{}, expire time.Duration) error
	KeyGet(ctx context.Context, key string) ([]byte, error)
	KeysGet(ctx context.Context, pattern string) ([]string, error)
	KeyDelete(ctx context.Context, key string) error
	KeyExist(ctx context.Context, key string) (bool, error)
	HashSet(ctx context.Context, key string, field string, value interface{}, expire time.Duration) error
	HashGet(ctx context.Context, key string, field string) ([]byte, error)
	HealthCheck(ctx context.Context) (string, error)
}
