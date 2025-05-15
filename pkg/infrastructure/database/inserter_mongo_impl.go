package database

import (
	"context"

	"property-service/pkg/errors"
	"property-service/pkg/helper/factory"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ Inserter[any] = (*InserterMongoImpl[any, any])(nil)

type InserterMongoImpl[DatabaseModel, DomainModel any] struct {
	log              log.Logger
	connector        Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection]
	factory          factory.Factory[DomainModel, DatabaseModel]
	collectionSuffix string
}

func NewMongoInserter[DatabaseModel, DomainModel any](
	log log.Logger,
	collectionSuffix string,
	factory factory.Factory[DomainModel, DatabaseModel],
	connector Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
) *InserterMongoImpl[DatabaseModel, DomainModel] {
	return &InserterMongoImpl[DatabaseModel, DomainModel]{
		log:              log,
		factory:          factory,
		connector:        connector,
		collectionSuffix: collectionSuffix,
	}
}

// InsertOne: This will insert one document of type DomainModel and return the id string or a error if one is returned.
func (imi *InserterMongoImpl[DatabaseModel, DomainModel]) InsertOne(
	c context.Context, server string, data DomainModel,
) (string, error) {
	// Check if the collection exists in the map of collections
	collection, collectionErr := imi.connector.GetCollection(server + imi.collectionSuffix)
	if collectionErr != nil {
		imi.log.Error("Collection not found %s", server+imi.collectionSuffix)
		return "", errors.ErrCollectionNotFound
	}

	// Map the document from domain to database.
	database, mappingErr := imi.factory.ToDatabase(data)
	if mappingErr != nil {
		imi.log.Error("Failed To Map Document: %+v", data)
		return "", errors.NewInternalError(mappingErr)
	}

	// Insert the document into the collection.
	res, insertErr := collection.InsertOne(c, database)
	if insertErr != nil {
		imi.log.Error("Error While inserting %+v", server+imi.collectionSuffix)
		return "", insertErr
	}
	// Debug log the output.
	imi.log.Debug("Collection: %s Successfully Inserted into Document : %+v with id %+v",
		server+imi.collectionSuffix, data, res.InsertedID)
	// Return the resulting ID.
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
