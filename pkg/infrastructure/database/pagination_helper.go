package database

import "go.mongodb.org/mongo-driver/bson"

// PaginationHelper defines a contract for building a paginated Atlas Search filter.
type PaginationHelper interface {
	// PaginationHelper builds an Atlas Search filter including pagination settings.
	TextPaginationHelper(
		index string,
		path string,
		value string,
		sort bson.D,
		search uint8,
		paginationToken string,
	) (bson.D, error)

	EqualsPaginationHelper(
		index string,
		path string,
		value string,
		sort bson.D,
		search uint8,
		paginationToken string,
	) (bson.D, error)
}
