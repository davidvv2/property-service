package service

import (
	"property-service/internal/properties/domain/owner"
	"property-service/internal/properties/domain/property"
	"property-service/pkg/configs"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type factories struct {
	Property property.Factory[uuid.UUID]
	Owner    owner.Factory[uuid.UUID]
}

func createFactories(
	_ log.Logger,
	v *validator.Validate,
	_ *configs.Config,
) factories {
	return factories{
		Property: property.MustNewFactory(
			property.FactoryConfig{
				SchemaVersion: 1,
			},
			v,
			database.NewStringID,
			database.IDToString,
			property.MapModelToProperty,
			database.StringToID,
			property.MapPropertyToModel,
		),
		Owner: owner.MustNewFactory(
			owner.FactoryConfig{
				SchemaVersion: 1,
			},
			v,
			database.NewStringID,
			database.IDToString,
			owner.MapModelToOwner,
			database.StringToID,
			owner.MapOwnerToModel,
		),
	}
}
