package database

import (
	"context"
)

type Grouper[Pipeline any, DomainModel any] interface {
	Aggregate(c context.Context, server string, params Pipeline) (Iterator[DomainModel], error)
}
