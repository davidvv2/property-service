package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"property-service/internal/properties/domain/property"
)

// CachedPropertyRepository is a decorator for property.Repository that adds Redis caching.
// It wraps a primary repository implementation (for example, our Mongo adapter) and uses
// a Redis adapter to cache read operations. Write operations update the primary store
// and then invalidate the cache.
type CachedPropertyRepository struct {
	// baseRepo is the primary repository (Mongo adapter).
	baseRepo property.Repository
	// redisAdapter is used for caching property data in Redis.
	redisAdapter *RedisCacheAdapter
	// cacheExpire defines how long cached entries are valid.
	cacheExpire time.Duration
}

// NewCachedPropertyRepository creates a new CachedPropertyRepository.
func NewCachedPropertyRepository(base property.Repository, redisAdapter *RedisCacheAdapter, expire time.Duration) *CachedPropertyRepository {
	return &CachedPropertyRepository{
		baseRepo:     base,
		redisAdapter: redisAdapter,
		cacheExpire:  expire,
	}
}

// generateCacheKey produces a key for caching single property items.
func generateCacheKey(operation, server, identifier string) string {
	return fmt.Sprintf("%s:%s:%s", operation, server, identifier)
}

// New creates a new property using the base repository.
// For new properties, we can also cache the created object.
func (c *CachedPropertyRepository) New(ctx context.Context, server string, parms property.NewPropertyParams) (*property.Property, error) {
	// Create the property using the Mongo adapter.
	newProp, err := c.baseRepo.New(ctx, server, parms)
	if err != nil {
		return nil, err
	}
	// Cache the newly created property.
	key := generateCacheKey("get", server, newProp.ID)
	data, err := json.Marshal(newProp)
	if err == nil {
		// We ignore cache errors so that caching issues do not block the main workflow.
		_ = c.redisAdapter.cacher.KeySet(ctx, key, data, c.cacheExpire)
	}

	return newProp, nil
}

// Delete removes a property using the base repository and invalidates its cached entry.
func (c *CachedPropertyRepository) Delete(ctx context.Context, server string, ID string) error {
	// Delete the property from Mongo.
	if err := c.baseRepo.Delete(ctx, server, ID); err != nil {
		return err
	}
	// Invalidate the cached property.
	key := generateCacheKey("get", server, ID)
	// Assuming our redis.Cacher supports a KeyDelete method; otherwise, adjust accordingly.
	_ = c.redisAdapter.cacher.KeyDelete(ctx, key)
	return nil
}

// Get retrieves a property: first it tries to fetch it from Redis cache,
// and if the cache is empty (or data is corrupted), it retrieves from Mongo and caches the result.
func (c *CachedPropertyRepository) Get(ctx context.Context, server string, ID string) (*property.Property, error) {
	key := generateCacheKey("get", server, ID)

	// Try to retrieve the property from Redis.
	cachedData, err := c.redisAdapter.cacher.KeyGet(ctx, key)
	if err == nil && cachedData != nil {
		var prop property.Property
		if err := json.Unmarshal(cachedData, &prop); err == nil {
			// Successfully retrieved the property from cache.
			return &prop, nil
		}
		// If unmarshalling fails, we fall back to Mongo.
	}

	// Retrieve the property from the Mongo repository.
	prop, err := c.baseRepo.Get(ctx, server, ID)
	if err != nil {
		return nil, err
	}

	// Cache the retrieved property for future calls.
	data, err := json.Marshal(prop)
	if err == nil {
		_ = c.redisAdapter.cacher.KeySet(ctx, key, data, c.cacheExpire)
	}
	return prop, nil
}

// Update uses the base repository to update a property and then invalidates its cache.
func (c *CachedPropertyRepository) Update(ctx context.Context, server string, id string, params property.UpdatePropertyParams) error {
	// Update the property in Mongo.
	err := c.baseRepo.Update(ctx, server, id, params)
	if err != nil {
		return err
	}
	// Invalidate the cache entry for the updated property.
	key := generateCacheKey("get", server, id)
	_ = c.redisAdapter.cacher.KeyDelete(ctx, key)
	return nil
}

// ListByCategory retrieves properties that match a given category,
// first checking the Redis cache and then falling back to Mongo if needed.
func (c *CachedPropertyRepository) ListByCategory(
	ctx context.Context,
	server string,
	category string,
	sort uint8,
	limit uint16,
	paginationToken string,
	search uint8,
) ([]property.Property, error) {
	// Construct a cache key that uniquely identifies the query.
	key := fmt.Sprintf("list:%s:%s:%s:%d:%d", server, category, paginationToken, sort, limit)

	// Attempt to get the cached list from Redis.
	cachedData, err := c.redisAdapter.cacher.KeyGet(ctx, key)
	if err == nil && cachedData != nil {
		var props []property.Property
		if err := json.Unmarshal(cachedData, &props); err == nil {
			// Found and unmarshalled the cached data successfully.
			return props, nil
		}
		// If unmarshalling fails, we ignore the cache and query Mongo.
	}

	// Retrieve the list from the primary repository (Mongo).
	props, err := c.baseRepo.ListByCategory(ctx, server, category, sort, limit, paginationToken, search)
	if err != nil {
		return nil, err
	}

	// Cache the result for future queries.
	data, err := json.Marshal(props)
	if err == nil {
		_ = c.redisAdapter.cacher.KeySet(ctx, key, data, c.cacheExpire)
	}
	return props, nil
}
