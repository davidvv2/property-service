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
	log              log.Logger
	factory          factory.Factory[DomainModel, DatabaseModel]
	connector        Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	collectionSuffix string
}

func NewMongoGrouper[DomainModel, DatabaseModel any](
	log log.Logger,
	factory factory.Factory[DomainModel, DatabaseModel],
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	collectionSuffix string,
) *GrouperMongoImpl[mongo.Pipeline, DomainModel, DatabaseModel] {
	return &GrouperMongoImpl[mongo.Pipeline, DomainModel, DatabaseModel]{
		log:              log,
		factory:          factory,
		connector:        connector,
		collectionSuffix: collectionSuffix,
	}
}

func (
	gmi *GrouperMongoImpl[Pipeline, DomainModel, DatabaseModel],
) Aggregate(
	c context.Context, server string, pipeline Pipeline,
) (Iterator[DomainModel], error) {
	collection, collectionErr := gmi.connector.GetCollection(server + gmi.collectionSuffix)
	if collectionErr != nil {
		return nil, errors.ErrCollectionNotFound
	}
	cursor, err := collection.Aggregate(c, pipeline)
	return NewMongoIterator[DomainModel, DatabaseModel](cursor, gmi.factory), err
}
