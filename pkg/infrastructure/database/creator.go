package database

import (
	"context"
)

var _ Creator = (*CreatorMongoImpl)(nil)

type Creator interface {
	CreateCollection(c context.Context, name string) error
	CreateIndex(c context.Context, collection, key, index string) (string, error)
}
