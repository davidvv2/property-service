package database

import (
	"context"
	"sync"

	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/helper/factory"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const BatchSize = 50

var _ Finder[bson.M, any] = (*FinderMongoImpl[bson.M, any, any])(nil)

type FinderMongoImpl[
	Filter bson.M,
	DomainModel, DatabaseModel any,
] struct {
	log       log.Logger
	Connector Connector[
		mongo.Client,
		mongo.ClientEncryption,
		mongo.Collection,
	]

	factory factory.Factory[DomainModel, DatabaseModel]

	findOneOptions *options.FindOneOptions
	findOptions    *options.FindOptions

	helper QueryHelper[
		Query[string],
		TimeQuery[string],
		options.FindOptions, bson.M,
	]
	collection string
}

func NewMongoFinder[DomainModel, DatabaseModel any](
	log log.Logger,
	collection string,
	factory factory.Factory[DomainModel, DatabaseModel],
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	findOneOptions *options.FindOneOptions,
	findOptions *options.FindOptions,
) *FinderMongoImpl[bson.M, DomainModel, DatabaseModel] {
	return &FinderMongoImpl[bson.M, DomainModel, DatabaseModel]{
		log:            log,
		factory:        factory,
		Connector:      connector,
		collection:     collection,
		findOneOptions: findOneOptions,
		findOptions:    findOptions,
		helper:         NewMongoQueryHelper(),
	}
}

func (fmi *FinderMongoImpl[
	Filter, DomainModel, DatabaseModel],
) FindOne(c context.Context, filter Filter) (*DomainModel, error) {
	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		return nil, errors.ErrCollectionNotFound
	}
	// Query the database and decode the result.
	res := collection.FindOne(c, filter, fmi.findOneOptions)
	var result DatabaseModel
	resultErr := res.Decode(&result)
	if resultErr != nil {
		return nil, errors.NewInternalError(resultErr)
	}
	// Map the result.
	return fmi.factory.ToDomain(result)
}

func (fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) Find(
	c context.Context, filter Filter,
) (<-chan *DomainModel, <-chan error, <-chan bool) {
	results := make(chan *DomainModel, BatchSize)
	errChan := make(chan error, BatchSize)
	doneChan := make(chan bool, 1)

	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		errChan <- errors.ErrCollectionNotFound
		return nil, errChan, doneChan
	}
	res, err := collection.Find(c, filter, fmi.findOptions)
	if err != nil {
		errChan <- errors.NewDatabaseError(err)
		return nil, errChan, doneChan
	}
	go fmi.itr(
		c,
		NewMongoIterator[DomainModel, DatabaseModel](res, fmi.factory),
		errChan, doneChan, results,
	)
	return results, errChan, doneChan
}

func (fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) FindByID(c context.Context, id string) (*DomainModel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewRepositoryError(
			err,
			codes.Internal,
		)
	}
	return fmi.FindOne(c, Filter(bson.M{"_id": oid}))
}

func (
	fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) find(
	c context.Context, filter Filter,
	results chan *DomainModel, errChan chan error, wg *sync.WaitGroup,
) {
	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		errChan <- errors.ErrCollectionNotFound
		return
	}
	res, err := collection.Find(c, filter, fmi.findOptions)
	if err != nil {
		errChan <- errors.NewDatabaseError(err)
		return
	}
	go func() {
		for res.Next(c) {
			var result DatabaseModel
			decodeErr := res.Decode(&result)
			if decodeErr != nil {
				errChan <- decodeErr
				continue
			}
			domainModel, conversionErr := fmi.factory.ToDomain(result)
			if conversionErr != nil {
				errChan <- conversionErr
				continue
			}
			results <- domainModel
		}
		wg.Done()
	}()
}

func (
	fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) Count(
	c context.Context, filter Filter,
) (int64, error) {
	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		fmi.log.Error("Collection: %s Operation:Count Error:%+v", collection.Name(), err, errors.ErrCollectionNotFound)
		return 0, errors.ErrCollectionNotFound
	}
	count, err := collection.CountDocuments(c, filter)
	if err != nil {
		fmi.log.Error("Collection: %s Operation:Count Error:%+v", collection.Name(), err)
		return 0, errors.NewInternalError(err)
	}
	return count, err
}

func (
	fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) DocumentExists(
	c context.Context, id string,
) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, errors.NewRepositoryError(
			err,
			codes.Internal,
		)
	}
	return fmi.Count(c, Filter(bson.M{"_id": oid}))
}

func (fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) Query(
	c context.Context, filter Filter, query Query[string],
) (<-chan *DomainModel, <-chan error, <-chan bool) {
	results, errChan, doneChan := fmi.createChannel(BatchSize)

	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		errChan <- errors.ErrCollectionNotFound
		return nil, errChan, doneChan
	}
	opt := fmi.options(query.Cursor, query.Limit, query.Order, query.Skip, query.Sort)
	res, err := collection.Find(c, filter, &opt)
	if err != nil {
		errChan <- errors.NewDatabaseError(err)
		return nil, errChan, doneChan
	}
	go fmi.itr(
		c,
		NewMongoIterator[DomainModel, DatabaseModel](res, fmi.factory),
		errChan, doneChan, results,
	)
	return results, errChan, doneChan
}

func (fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) TimeQuery(
	c context.Context, filter Filter, query TimeQuery[string],
) (<-chan *DomainModel, int64, <-chan error, <-chan bool) {
	results, errChan, doneChan := fmi.createChannel(BatchSize)

	collection, err := fmi.Connector.GetCollection(fmi.collection)
	if err != nil {
		errChan <- errors.ErrCollectionNotFound
		return nil, 0, errChan, doneChan
	}
	opt := fmi.options(query.Cursor, query.Limit, query.Order, query.Skip, query.Sort)
	res, err := collection.Find(c, filter, &opt)
	if err != nil {
		errChan <- errors.NewDatabaseError(err)
		return nil, 0, errChan, doneChan
	}
	go fmi.itr(
		c,
		NewMongoIterator[DomainModel, DatabaseModel](res, fmi.factory),
		errChan, doneChan, results,
	)
	return results, 0, errChan, doneChan
}

func (fmi *FinderMongoImpl[Filter, DomainModel, DatabaseModel],
) itr(
	c context.Context,
	iterator *IteratorMongoImpl[DomainModel, DatabaseModel],
	errChan chan error,
	doneChan chan bool,
	results chan *DomainModel,
) {
	for hasNext, domainModel, itrErr := iterator.GetNext(c); hasNext; {
		if itrErr != nil {
			errChan <- itrErr
			return
		}
		results <- domainModel
	}
	doneChan <- true
	fmi.closeChannel(results, errChan, doneChan)
}

func (fmi *FinderMongoImpl[
	Filter,
	DomainModel,
	DatabaseModel,
]) options(
	cursor string,
	limit, order, skip int64,
	sort string,
) options.FindOptions {
	primitiveCursor, _ := primitive.ObjectIDFromHex(cursor)
	if primitiveCursor == primitive.NilObjectID {
		return fmi.helper.optionsHelper(limit, order, skip, sort, false)
	}
	return fmi.helper.optionsHelper(limit, order, skip, sort, true)
}

func (fmi *FinderMongoImpl[
	Filter,
	DomainModel,
	DatabaseModel,
]) createChannel(
	length int,
) (
	chan *DomainModel,
	chan error,
	chan bool,
) {
	return make(chan *DomainModel, length),
		make(chan error, length),
		make(chan bool, 1)
}

func (fmi *FinderMongoImpl[
	Filter,
	DomainModel,
	DatabaseModel,
]) closeChannel(
	results chan<- *DomainModel,
	errChan chan<- error,
	doneChan chan<- bool,
) {
	close(doneChan)
	close(results)
	close(errChan)
}
