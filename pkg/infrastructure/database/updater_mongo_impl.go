package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/helper/factory"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ Updater[bson.M, bson.M, any] = (*UpdaterMongoImpl[bson.M, bson.M, any, any])(nil)

type UpdaterMongoImpl[
	Filter,
	Partial bson.M,
	DomainModel,
	DatabaseModel any,
] struct {
	log     log.Logger
	factory factory.Factory[
		DomainModel,
		DatabaseModel,
	]
	connector Connector[
		mongo.Client,
		mongo.ClientEncryption,
		mongo.Collection,
	]
	collection string
}

func NewMongoUpdater[DomainModel, DatabaseModel any](
	log log.Logger,
	factory factory.Factory[DomainModel, DatabaseModel],
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	collection string,
) *UpdaterMongoImpl[bson.M, bson.M, DomainModel, DatabaseModel] {
	return &UpdaterMongoImpl[
		bson.M, bson.M, DomainModel, DatabaseModel]{
		log:        log,
		factory:    factory,
		connector:  connector,
		collection: collection,
	}
}

func (fmi *UpdaterMongoImpl[Filter, Partial, DomainModel, DatabaseModel]) UpdateOne(
	c context.Context, filter Filter, data Partial,
) error {
	fmi.log.Info("filter %+v data %+v", bson.M(filter), data)

	collection, err := fmi.connector.GetCollection(
		fmi.collection,
	)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(
		c,
		bson.M(filter),
		data,
	)
	return err
}

func (fmi *UpdaterMongoImpl[Filter, Partial, DomainModel, DatabaseModel]) UpdateOneByID(
	c context.Context, id string, data Partial,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmi.log.Error("ID from primitive err  %+v", err)
		return errors.NewRepositoryError(
			err,
			codes.Internal,
		)
	}
	return fmi.UpdateOne(
		c,
		Filter(bson.M{
			"_id": oid,
		}),
		data,
	)
}

func (fmi *UpdaterMongoImpl[
	Filter, Partial, DomainModel, DatabaseModel,
]) ReplaceOne(
	c context.Context, filter Filter, data DomainModel,
) (string, error) {
	collection, err := fmi.connector.GetCollection(fmi.collection)
	if err != nil {
		return "", err
	}
	res, err := collection.ReplaceOne(
		c,
		filter,
		data,
	)
	var updated bool
	switch {
	case res.MatchedCount > 0:
		updated = true
	case res.ModifiedCount > 0:
		updated = true
	case res.UpsertedCount > 0:
		updated = true
	}
	if !updated {
		return "", errors.ErrNoUpdate
	}
	return res.UpsertedID.(primitive.ObjectID).Hex(), err
}

func (fmi *UpdaterMongoImpl[
	Filter,
	Partial,
	DomainModel,
	DatabaseModel,
]) UpdateAndFind(
	c context.Context,

	filter Filter,
	data Partial,
) (*DomainModel, error) {
	collection, err := fmi.connector.GetCollection(
		fmi.collection,
	)
	if err != nil {
		return nil, err
	}
	// Find and update the model in the database.
	result := collection.FindOneAndUpdate(
		c,
		filter,
		data,
	)

	// Decode the database result.
	var dbModel DatabaseModel
	findErr := result.Decode(&dbModel)
	if findErr != nil {
		return nil, findErr
	}

	// Return the domain model.
	return fmi.factory.ToDomain(
		dbModel,
	)
}
