package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ Remover[bson.M] = (*RemoverMongoImpl[bson.M])(nil)

type RemoverMongoImpl[Filter bson.M] struct {
	log        log.Logger
	connector  Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	collection string
}

func NewMongoRemover(
	log log.Logger,
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	collection string,
) *RemoverMongoImpl[bson.M] {
	return &RemoverMongoImpl[bson.M]{
		log:        log,
		connector:  connector,
		collection: collection,
	}
}

func (rmi *RemoverMongoImpl[Filter]) DeleteOne(
	c context.Context, filter Filter,
) (int64, error) {
	// Check if the collection exists in the map of collections
	collection, collectionErr := rmi.connector.GetCollection(rmi.collection)
	if collectionErr != nil {
		rmi.log.Error("Collection not found %s", rmi.collection)
		return 0, errors.ErrCollectionNotFound
	}
	// Delete the item by the filer.
	deleteResult, deleteErr := collection.DeleteOne(c, filter)
	if deleteResult.DeletedCount == 0 {
		return 0, errors.NewInternalError(errors.New("delete count was 0"))
	}
	return deleteResult.DeletedCount, deleteErr
}

func (rmi *RemoverMongoImpl[Filter]) DeleteOneByID(
	c context.Context, id string,
) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, errors.NewRepositoryError(
			err,
			codes.Internal,
		)
	}
	return rmi.DeleteOne(c, Filter(bson.M{"_id": oid}))
}
