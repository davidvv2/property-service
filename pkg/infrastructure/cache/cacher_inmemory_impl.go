package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

type inMemoryCacherImpl struct {
	values map[string][]byte
	hashes map[string]map[string][]byte
	mutex  sync.RWMutex
}

// NewMemoryCacher creates a new in-memory cache instance.
func NewMemoryCacher() Cacher {
	return &inMemoryCacherImpl{
		values: make(map[string][]byte),
		hashes: make(map[string]map[string][]byte),
		mutex:  sync.RWMutex{},
	}
}

// KeyExist : checks to see if a key exists.
func (c *inMemoryCacherImpl) KeyExist(_ context.Context, key string) (bool, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.values[key]
	return ok, nil
}

func (c *inMemoryCacherImpl) KeyGet(_ context.Context, key string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.values[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (c *inMemoryCacherImpl) KeySet(_ context.Context, key string, value interface{}, expire time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("can not cast")
	}
	c.values[key] = bytes
	callBack := time.NewTimer(expire)
	go func(c *inMemoryCacherImpl) {
		<-callBack.C
		c.hashes[key] = nil
	}(c)
	return nil
}

func (c *inMemoryCacherImpl) HashGet(_ context.Context, key string, field string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.hashes[key][field]
	if !ok {
		return nil, errors.New("field not found")
	}
	return val, nil
}

func (c *inMemoryCacherImpl) HashSet(
	_ context.Context, key string, field string, value interface{}, expire time.Duration,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.hashes[key] == nil {
		c.hashes[key] = make(map[string][]byte)
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("can not cast")
	}

	c.hashes[key][field] = bytes
	callBack := time.NewTimer(expire)
	go func(c *inMemoryCacherImpl) {
		<-callBack.C
		c.hashes[key][field] = nil
	}(c)

	return nil
}

func (c *inMemoryCacherImpl) HealthCheck(_ context.Context) (string, error) {
	return "OK", nil
}
