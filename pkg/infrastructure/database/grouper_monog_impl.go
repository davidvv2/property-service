package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/helper/factory"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ Grouper[mongo.Pipeline, any] = (*GrouperMongoImpl[
	mongo.Pipeline, any, any,
])(nil)

type GrouperMongoImpl[Pipeline mongo.Pipeline, DomainModel, DatabaseModel any] struct {
	log        log.Logger
	factory    factory.Factory[DomainModel, DatabaseModel]
	connector  Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	collection string
}

func NewMongoGrouper[DomainModel, DatabaseModel any](
	log log.Logger,
	factory factory.Factory[DomainModel, DatabaseModel],
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	collection string,
) *GrouperMongoImpl[mongo.Pipeline, DomainModel, DatabaseModel] {
	return &GrouperMongoImpl[mongo.Pipeline, DomainModel, DatabaseModel]{
		log:        log,
		factory:    factory,
		connector:  connector,
		collection: collection,
	}
}

func (
	gmi *GrouperMongoImpl[Pipeline, DomainModel, DatabaseModel],
) Aggregate(
	c context.Context, pipeline Pipeline,
) (Iterator[DomainModel], error) {
	collection, collectionErr := gmi.connector.GetCollection(gmi.collection)
	if collectionErr != nil {
		return nil, errors.ErrCollectionNotFound
	}
	cursor, err := collection.Aggregate(c, pipeline)
	return NewMongoIterator[DomainModel, DatabaseModel](cursor, gmi.factory), err
}
