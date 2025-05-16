package service

import (
	"property-service/internal/properties/adapters"
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/domain/property"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

const _DatabaseName = "properties"

type repositories struct {
	PropertyRepository property.Repository
	OwnerRepository    owner.Repository
}

func createRepositories(
	l log.Logger,
	config *configs.Config,
	factory factories,
	v *validator.Validate,
) repositories {
	connector := database.NewMongoConnector(
		l,
		config.Database,
		_DatabaseName,
	)

	prop := createProperty(
		l,
		factory,
		v,
		connector,
		config.Database,
	)

	propRepo := adapters.NewMongoPropertyRepository(
		l,
		prop.FinderInsterterUpdaterRemover,
		factory.Property,
		prop.Aggregator,
	)

	owner := createOwner(
		l,
		factory,
		v,
		connector,
		config.Database,
	)

	ownerRepo := adapters.NewMongoOwnerRepository(
		l,
		owner.FinderInsterterUpdaterRemover,
		factory.Owner,
	)
	return repositories{
		PropertyRepository: propRepo,
		OwnerRepository:    ownerRepo,
	}
}
