package adapters

import (
	"context"
	"encoding/json"
	"time"

	"property-service/internal/properties/domain/property" // adjust the import if needed
	redis "property-service/pkg/infrastructure/cache"      // redis cacher package
	"property-service/pkg/infrastructure/log"
)

// RedisCacheAdapter implements caching operations for property data using Redis.
type RedisCacheAdapter struct {
	cacher redis.Cacher
	log    log.Logger // Logger for logging operations, if needed
}

// NewRedisCacheAdapter creates a new RedisCacheAdapter with the given redis cacher.
func NewRedisCacheAdapter(cacher redis.Cacher, log log.Logger) *RedisCacheAdapter {
	return &RedisCacheAdapter{
		cacher: cacher,
		log:    log,
	}
}

// SaveProperty caches a property's data with the given expiration.
func (r *RedisCacheAdapter) SaveProperty(ctx context.Context, p property.Property, expire time.Duration) error {
	// Marshal the property into JSON
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	// Create a cache key based on the property ID
	cacheKey := "property:" + p.ID
	return r.cacher.KeySet(ctx, cacheKey, data, expire)
}

// GetProperty retrieves a cached property using its ID.
func (r *RedisCacheAdapter) GetProperty(ctx context.Context, id string) (*property.Property, error) {
	cacheKey := "property:" + id
	data, err := r.cacher.KeyGet(ctx, cacheKey)
	if err != nil {
		return nil, err
	}
	if data == nil {
		// Cache miss; the property is not in cache.
		return nil, nil
	}
	var prop property.Property
	err = json.Unmarshal(data, &prop)
	if err != nil {
		r.log.ErrorWithFields("Failed to unmarshal property data from cache", log.Fields{
			"error": err,
		})
		return nil, err
	}
	return &prop, nil
}

// ListProperties retrieves a list of cached properties.
func (r *RedisCacheAdapter) ListProperties(ctx context.Context, server string) ([]property.Property, error) {
	// This method would typically involve scanning keys in Redis.
	// For simplicity, we assume a specific pattern for property keys.
	prefix := "property:"
	keys, err := r.cacher.KeysGet(ctx, prefix+"*")
	if err != nil {
		return nil, err
	}

	var properties []property.Property
	for _, key := range keys {
		data, err := r.cacher.KeyGet(ctx, string(key))
		if err != nil {
			r.log.DebugWithFields("Failed to get property data from cache", log.Fields{
				"key":   key,
				"error": err,
			})
			continue // Skip this key if there's an error
		}
		if data == nil {
			r.log.DebugWithFields("No data found for property key", log.Fields{
				"key": key,
			})
			continue // Skip if no data found
		}
		var prop property.Property
		if err := json.Unmarshal(data, &prop); err == nil {
			properties = append(properties, prop)
		}
	}
	return properties, nil
}
