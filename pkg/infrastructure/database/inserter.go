package database

import (
	"context"
)

type Inserter[DomainModel any] interface {
	InsertOne(c context.Context, data DomainModel) (string, error)
}
