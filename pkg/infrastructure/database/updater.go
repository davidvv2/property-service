package database

import (
	"context"
)

type Updater[Filter, Partial any, DomainModel any] interface {
	UpdateOne(
		c context.Context,
		filter Filter,
		data Partial,
	) error

	UpdateOneByID(
		c context.Context,
		id string,
		data Partial,
	) error

	ReplaceOne(
		c context.Context,
		filter Filter,
		data DomainModel,
	) (string, error)

	UpdateAndFind(
		c context.Context,
		filter Filter,
		data Partial,
	) (*DomainModel, error)
}
