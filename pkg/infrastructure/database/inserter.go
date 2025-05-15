package database

import (
	"context"
)

type Inserter[DomainModel any] interface {
	InsertOne(c context.Context, server string, data DomainModel) (string, error)
}
