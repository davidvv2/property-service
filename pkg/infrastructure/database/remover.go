package database

import (
	"context"
)

type Remover[Filter any] interface {
	DeleteOne(c context.Context, filter Filter) (int64, error)
	DeleteOneByID(c context.Context, ID string) (int64, error)
}
