package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/helper/factory"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ Iterator[any] = (*IteratorMongoImpl[any, any])(nil)

type IteratorMongoImpl[DomainModel, DatabaseModel any] struct {
	cursor  *mongo.Cursor
	factory factory.Factory[DomainModel, DatabaseModel]
}

func NewMongoIterator[DomainModel, DatabaseModel any](
	cursor *mongo.Cursor,
	factory factory.Factory[DomainModel, DatabaseModel],
) *IteratorMongoImpl[DomainModel, DatabaseModel] {
	return &IteratorMongoImpl[DomainModel, DatabaseModel]{
		cursor:  cursor,
		factory: factory,
	}
}

func (imi *IteratorMongoImpl[DomainModel, DatabaseModel]) GetAll(c context.Context) (*[]DomainModel, error) {
	var result []DatabaseModel
	allErr := imi.cursor.All(c, &result)
	if allErr != nil {
		return nil, errors.NewInternalError(allErr)
	}
	return imi.mapArray(len(result), result)
}

func (imi *IteratorMongoImpl[DomainModel, DatabaseModel]) GetNext(c context.Context) (bool, *DomainModel, error) {
	hasNext := imi.cursor.Next(c)
	if hasNext {
		var result DatabaseModel
		decodeErr := imi.cursor.Decode(&result)
		if decodeErr != nil {
			return hasNext, nil, errors.NewInternalError(decodeErr)
		}
		res, err := imi.factory.ToDomain(result)
		if err != nil {
			return hasNext, nil, err
		}
		return hasNext, res, nil
	}
	return hasNext, nil, errors.NewDatabaseError(errors.New("end of iterator"))
}

func (
	imi *IteratorMongoImpl[DomainModel, DatabaseModel],
) GetNextBatch(
	c context.Context,
) (*[]DomainModel, error) {
	var result []DatabaseModel
	if !imi.cursor.Next(c) {
		return nil, errors.NewInternalError(errors.New("could not iter"))
	}
	decodeErr := imi.cursor.Decode(&result)
	if decodeErr != nil {
		return nil, errors.NewInternalError(decodeErr)
	}
	return imi.mapArray(len(result), result)
}

func (imi *IteratorMongoImpl[DomainModel, DatabaseModel]) GetBatchLength() int {
	return imi.cursor.RemainingBatchLength()
}

func (imi *IteratorMongoImpl[DomainModel, DatabaseModel]) Close(c context.Context) error {
	return imi.cursor.Close(c)
}

func (imi *IteratorMongoImpl[DomainModel, DatabaseModel]) mapArray(
	resultLen int, result []DatabaseModel,
) (*[]DomainModel, error) {
	domainArray := make([]DomainModel, resultLen)
	for i := 0; i < resultLen; i++ {
		res, err := imi.factory.ToDomain(result[i])
		if err != nil {
			return nil, err
		}
		domainArray[i] = *res
	}
	return &domainArray, nil
}
