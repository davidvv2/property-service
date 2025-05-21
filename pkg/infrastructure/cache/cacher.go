// Package cache : Handles all the Cache related functions.
package cache

import (
	"context"
	"time"
)

// Cacher : used to cache items as either a key value pair or values in a hash table.
type Cacher interface {
	/*******************
	 * Key value pairs *
	 *******************/
	// Check to see if a given key has a value.
	KeyExist(c context.Context, key string) (bool, error)
	// Get a value stored in a given key.
	KeyGet(c context.Context, key string) ([]byte, error)
	// Set a value of a given key.
	KeySet(c context.Context, key string, value interface{}, expire time.Duration) error

	/***************
	 * Hash table  *
	 ***************/
	// Get a value from a given key.
	HashGet(c context.Context, key string, field string) ([]byte, error)
	// Set a value of a given key in a hash table.
	HashSet(c context.Context, key string, field string, value interface{}, expire time.Duration) error

	/*****************
	 * Health Check  *
	 *****************/
	// Check the health of the caching object.
	HealthCheck(c context.Context) (string, error)
}
