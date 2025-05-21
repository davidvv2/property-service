package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreatorMongoImpl struct {
	log       log.Logger
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
}

func NewMongoCreator(
	log log.Logger,
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
) *CreatorMongoImpl {
	return &CreatorMongoImpl{
		log:       log,
		connector: connector,
	}
}

func (cmi *CreatorMongoImpl) CreateCollection(c context.Context, name string) error {
	err := cmi.connector.getClient().Database(cmi.connector.GetDatabaseName()).
		CreateCollection(c, name, &options.CreateCollectionOptions{})
	if err != nil {
		return errors.NewInternalError(err)
	}
	return nil
}

func (cmi *CreatorMongoImpl) CreateIndex(c context.Context, collection, key, index string) (string, error) {
	coll, err := cmi.connector.GetCollection(collection)
	if err != nil {
		return "", errors.NewDatabaseError(err)
	}
	res, err := coll.Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{{Key: key, Value: index}},
	})
	if err != nil {
		return "", errors.NewDatabaseError(err)
	}
	return res, nil
}
