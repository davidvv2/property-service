package database

import (
	"context"
)

type Updater[Filter, Partial any, DomainModel any] interface {
	UpdateOne(
		c context.Context,
		server string,
		filter Filter,
		data Partial,
	) error

	UpdateOneByID(
		c context.Context,
		server string,
		id string,
		data Partial,
	) error

	ReplaceOne(
		c context.Context,
		server string,
		filter Filter,
		data DomainModel,
	) (string, error)

	UpdateAndFind(
		c context.Context,
		server string,
		filter Filter,
		data Partial,
	) (*DomainModel, error)
}
