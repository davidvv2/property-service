package database

import (
	"context"
)

type Remover[Filter any] interface {
	DeleteOne(c context.Context, server string, filter Filter) (int64, error)
	DeleteOneByID(c context.Context, server string, ID string) (int64, error)
}
