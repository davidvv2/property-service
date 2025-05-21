package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/log"

	"github.com/go-redis/redis/v8"
)

// redisCacherImpl : implements a redis version of the cacher interface.
type redisCacherImpl struct {
	// Client : the redis client object
	Client *redis.Client
	log    log.Logger
}

// NewRedisCacher : Connects to a redis instance.
func NewRedisCacher(l log.Logger, config *configs.CachingStruct) Cacher {
	//nolint:exhaustruct //do not need to provide all redis options
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	/////todo add ping test here and return error
	l.Info("Connected to Redis %s", redisClient.Ping(context.Background()))
	return &redisCacherImpl{Client: redisClient, log: l}
}

// Set a value of a given key.
func (rs redisCacherImpl) KeySet(c context.Context, key string, value interface{}, expire time.Duration) error {
	if err := rs.Client.Set(c, key, value, expire).Err(); err != nil {
		rs.log.Error("Error while trying to set value", err.Error())
		// print error and retry one more time then exit if fail.
		if err = rs.Client.Set(c, key, value, expire).Err(); err != nil {
			return fmt.Errorf("failed to set key: %w", err)
		}
	}
	return nil
}

// KeyGet :gets a value from a given key.
func (rs redisCacherImpl) KeyGet(ctx context.Context, key string) ([]byte, error) {
	cache, err := rs.Client.Get(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("failed to get key: %w", err)
	}

	return []byte(cache), nil
}

// KeyExist :checks to see if a key exists.
func (rs redisCacherImpl) KeyExist(ctx context.Context, key string) (bool, error) {
	response, err := rs.Client.Exists(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("failed to check key: %w", err)
	case response != 0:
		return true, nil
	}

	return false, nil
}

// Set :Sets a given value for a given key.
func (rs redisCacherImpl) HashSet(
	c context.Context, key string, field string, value interface{}, expire time.Duration,
) error {
	// Set the hash field.
	if err := rs.Client.HSet(c, key, field, value).Err(); err != nil {
		rs.log.Error("Error while trying to set hash field", log.Fields{"error": err.Error()})
		// Optionally retry setting the field before failing.
		if retryErr := rs.Client.HSet(c, key, field, value).Err(); retryErr != nil {
			return fmt.Errorf("failed to set hash key: %w", retryErr)
		}
	}
	// Set the expiration on the key separately.
	if err := rs.Client.Expire(c, key, expire).Err(); err != nil {
		return fmt.Errorf("failed to set expiration on hash key: %w", err)
	}
	return nil
}

// Get a value from a given key.
func (rs redisCacherImpl) HashGet(
	ctx context.Context, key string, field string) ([]byte, error,
) {
	cache, err := rs.Client.HGet(ctx, key, field).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("failed to get hash key: %w", err)
	}

	return []byte(cache), nil
}

/*****************
 * Health Check  *
 *****************/

// HealthCheck: Check the health of the caching object.
func (rs redisCacherImpl) HealthCheck(
	ctx context.Context,
) (string, error) {
	status, err := rs.Client.Ping(ctx).Result()
	return status, fmt.Errorf("failed to initialise health test: %w", err)
}
