package database

import "context"

type Session[SessionFunction any] interface {
	Execute(c context.Context, call SessionFunction) error
}
