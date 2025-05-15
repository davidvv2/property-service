package database

import "context"

type Iterator[DomainModel any] interface {
	GetAll(context.Context) (*[]DomainModel, error)
	GetNext(context.Context) (bool, *DomainModel, error)
	GetNextBatch(context.Context) (*[]DomainModel, error)
	GetBatchLength() int
	Close(context.Context) error
}
