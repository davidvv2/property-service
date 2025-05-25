package service

import (
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/domain/property"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collections constants.
const (
	_PROPERTY = "Property"
	_OWNER    = "Owner"
)

type Property struct {
	finder *database.FinderMongoImpl[
		primitive.M, property.Property, property.Model[primitive.ObjectID],
	]
	updater *database.UpdaterMongoImpl[
		primitive.M, primitive.M, property.Property, property.Model[primitive.ObjectID],
	]
	FinderUpdater database.FinderUpdater[
		bson.M, bson.M, property.Property,
	]
	Inserter                      *database.InserterMongoImpl[property.Model[primitive.ObjectID], property.Property]
	FinderInsterterUpdaterRemover database.FinderInserterUpdaterRemover[
		bson.M, bson.M, property.Property,
	]
	Aggregator database.Grouper[
		mongo.Pipeline, property.Property,
	]
}

func createProperty(
	l log.Logger,
	factory factories,
	v *validator.Validate,
	connector database.Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	config configs.DatabaseStruct,
) Property {
	// Finder
	propFinder := database.NewMongoFinder(
		l, _PROPERTY, factory.Property, connector,
		options.FindOne(), options.Find())
	// Updater
	propUpdater := database.NewMongoUpdater(
		l, factory.Property, connector, _PROPERTY,
	)
	// Inserter
	propInserter := database.NewMongoInserter(
		l, _PROPERTY, factory.Property, connector,
	)
	// FinderUpdater
	propFinderUpdater := database.NewMongoFinderUpdater(propFinder, propUpdater)

	// Remover
	propRemover := database.NewMongoRemover(l, connector, _PROPERTY)
	// FinderInserterUpdaterRemover
	propFinderInserterUpdaterRemover := database.NewMongoFinderInserterUpdaterRemover(
		propFinder, propInserter, propUpdater, propRemover,
	)

	// Aggregator
	propAggregator := database.NewMongoGrouper(
		l, factory.Property, connector, _PROPERTY,
	)

	return Property{
		finder:                        propFinder,
		updater:                       propUpdater,
		FinderUpdater:                 propFinderUpdater,
		Inserter:                      propInserter,
		FinderInsterterUpdaterRemover: propFinderInserterUpdaterRemover,
		Aggregator:                    propAggregator,
	}
}

type Owner struct {
	finder *database.FinderMongoImpl[
		primitive.M, owner.Owner, owner.Model[primitive.ObjectID],
	]
	updater *database.UpdaterMongoImpl[
		primitive.M, primitive.M, owner.Owner, owner.Model[primitive.ObjectID],
	]
	FinderUpdater database.FinderUpdater[
		bson.M, bson.M, owner.Owner,
	]
	Inserter                      *database.InserterMongoImpl[owner.Model[primitive.ObjectID], owner.Owner]
	FinderInsterterUpdaterRemover database.FinderInserterUpdaterRemover[
		bson.M, bson.M, owner.Owner,
	]
}

func createOwner(
	l log.Logger,
	factory factories,
	v *validator.Validate,
	connector database.Connector[mongo.Client, mongo.ClientEncryption, mongo.Collection],
	config configs.DatabaseStruct,
) Owner {
	// Finder
	ownerFinder := database.NewMongoFinder(
		l, _OWNER, factory.Owner, connector,
		options.FindOne(), options.Find())
	// Updater
	ownerUpdater := database.NewMongoUpdater(
		l, factory.Owner, connector, _OWNER,
	)
	// Inserter
	ownerInserter := database.NewMongoInserter(
		l, _OWNER, factory.Owner, connector,
	)
	// FinderUpdater
	ownerFinderUpdater := database.NewMongoFinderUpdater(ownerFinder, ownerUpdater)

	// Remover
	ownerRemover := database.NewMongoRemover(l, connector, _OWNER)
	// FinderInserterUpdaterRemover
	ownerFinderInserterUpdaterRemover := database.NewMongoFinderInserterUpdaterRemover(
		ownerFinder, ownerInserter, ownerUpdater, ownerRemover,
	)

	return Owner{
		finder:                        ownerFinder,
		updater:                       ownerUpdater,
		FinderUpdater:                 ownerFinderUpdater,
		Inserter:                      ownerInserter,
		FinderInsterterUpdaterRemover: ownerFinderInserterUpdaterRemover,
	}
}
