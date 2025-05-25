package redis

import (
	"context"
	"errors"
	"fmt"

	"time"

	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"

	"github.com/go-redis/redis/v8"
)

// customRedisCacherImpl implements a custom Redis version of the cacher interface.
type RedisCacherImpl struct {
	Client *redis.Client
	log    log.Logger
}

// NewCustomRedisCacher connects to a Redis instance and returns a new customRedisCacherImpl.
func NewRedisCacher(config *configs.CachingStruct, log log.Logger) *RedisCacherImpl {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return &RedisCacherImpl{Client: redisClient, log: log}
}

// KeySet sets a value for a given key with an expiration time.
func (rc *RedisCacherImpl) KeySet(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	if err := rc.Client.Set(ctx, key, value, expire).Err(); err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}
	return nil
}

// KeyGet retrieves a value from a given key.
func (rc *RedisCacherImpl) KeyGet(ctx context.Context, key string) ([]byte, error) {
	cache, err := rc.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get key: %w", err)
	}
	return []byte(cache), nil
}

// KeysGet retrieves all keys matching a given pattern.
func (rc *RedisCacherImpl) KeysGet(ctx context.Context, pattern string) ([]string, error) {
	keys, err := rc.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}
	var results []string
	for _, key := range keys {
		value, err := rc.Client.Get(ctx, key).Result()
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				return nil, fmt.Errorf("failed to get key %s: %w", key, err)
			}
			// If the key does not exist, we can skip it.
			continue
		}
		results = append(results, value)
	}
	return results, nil
}

// KeyDelete deletes a key from the Redis store.
func (rc *RedisCacherImpl) KeyDelete(ctx context.Context, key string) error {
	if err := rc.Client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}

// KeyExist checks if a key exists in the Redis store.
func (rc *RedisCacherImpl) KeyExist(ctx context.Context, key string) (bool, error) {
	response, err := rc.Client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check key: %w", err)
	}
	return response != 0, nil
}

// HashSet sets a field in a hash stored at a key with an expiration time.
func (rc *RedisCacherImpl) HashSet(ctx context.Context, key string, field string, value interface{}, expire time.Duration) error {
	if err := rc.Client.HSet(ctx, key, field, value).Err(); err != nil {
		return fmt.Errorf("failed to set hash: %w", err)
	}
	if err := rc.Client.Expire(ctx, key, expire).Err(); err != nil {
		return fmt.Errorf("failed to set expiration: %w", err)
	}
	return nil
}

// HashGet retrieves a field from a hash stored at a key.
func (rc *RedisCacherImpl) HashGet(ctx context.Context, key string, field string) ([]byte, error) {
	cache, err := rc.Client.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get hash: %w", err)
	}
	return []byte(cache), nil
}

// HealthCheck checks the health of the Redis connection.
func (rc *RedisCacherImpl) HealthCheck(ctx context.Context) (string, error) {
	if err := rc.Client.Ping(ctx).Err(); err != nil {
		return "", fmt.Errorf("failed to ping Redis: %w", err)
	}
	return "PONG", nil
}
