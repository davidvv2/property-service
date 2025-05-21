package property

import (
	"time"

	"property-service/pkg/errors"

	"github.com/go-playground/validator/v10"
)

var _ Factory[any] = (*FactoryImpl[any])(nil)

// Factory is a struct that creates and validates the model.
type FactoryImpl[databaseID comparable] struct {
	NewID             func() string
	mapToDomainFunc   func(databaseID) (string, error)
	mapToDatabaseFunc func(string) (databaseID, error)
	mapToDomain       func(mapper func(databaseID) (string, error), his Model[databaseID]) (*Property, error)
	mapToDatabase     func(mapper func(string) (databaseID, error), his Property) (*Model[databaseID], error)
	v                 *validator.Validate // validator used for validating the factory configuration
	fc                FactoryConfig       // configuration for the factory
}

// NewFactory creates a new Factory with the given configuration and returns an error if the configuration is invalid.
func NewFactory[databaseID comparable](
	fc FactoryConfig, v *validator.Validate, newID func() string,
	mappingFunc func(databaseID) (string, error),
	mapHistory func(mappingFunc func(databaseID) (string, error), databaseModel Model[databaseID]) (*Property, error),
	mapToDatabaseFunc func(string) (databaseID, error),
	mapToDatabase func(mappingFunc func(string) (databaseID, error), domainModel Property) (*Model[databaseID], error),
) (FactoryImpl[databaseID], error) {
	if err := fc.Validate(); err != nil {
		return FactoryImpl[databaseID]{}, errors.Join(err, errors.ErrInvalidConfigFactory)
	}
	return FactoryImpl[databaseID]{
		fc:                fc,
		v:                 v,
		NewID:             newID,
		mapToDomainFunc:   mappingFunc,
		mapToDomain:       mapHistory,
		mapToDatabase:     mapToDatabase,
		mapToDatabaseFunc: mapToDatabaseFunc,
	}, nil
}

// MustNewFactory creates a new Factory with the given configuration and panics if the configuration is invalid.
func MustNewFactory[databaseID comparable](
	fc FactoryConfig, v *validator.Validate, newID func() string,
	mappingFunc func(databaseID) (string, error),
	mapHistory func(mappingFunc func(databaseID) (string, error), databaseModel Model[databaseID]) (*Property, error),
	mapToDatabaseFunc func(string) (databaseID, error),
	mapToDatabase func(mappingFunc func(string) (databaseID, error), domainModel Property) (*Model[databaseID], error),
) FactoryImpl[databaseID] {
	f, err := NewFactory[databaseID](fc, v, newID, mappingFunc, mapHistory, mapToDatabaseFunc, mapToDatabase)
	if err != nil {
		panic(err)
	}
	return f
}

// Config returns the configuration for the factory.
func (fi FactoryImpl[databaseID]) Config() FactoryConfig {
	return fi.fc
}

func (fi FactoryImpl[databaseID]) validate(la *Property) error {
	return fi.v.Struct(la)
}

type NewPropertyParams struct {
	PropertyID    string    `validate:"required"`
	OwnerID       string    `validate:"required"`
	Category      string    `validate:"required"`
	Description   string    `validate:"required"`
	Title         string    `validate:"required"`
	Available     bool      `validate:"required"`
	AvailableDate time.Time `validate:"required"`
	Address       string    `validate:"required"`
	SaleType      uint8     `validate:"required"`
}

func (fi FactoryImpl[databaseID]) New(
	property NewPropertyParams,
) (*Property, error) {
	propertyModel := &Property{
		id:          property.PropertyID,
		ownerID:     property.OwnerID,
		category:    property.Category,
		description: property.Description,
		title:       property.Title,
		metadata: Metadata{
			createdAt: time.Now(),
			updatedAt: time.Time{},
		},
		available:     property.Available,
		availableDate: property.AvailableDate,
		address:       property.Address,
		saleType:      property.SaleType,
	}
	return propertyModel, fi.validate(propertyModel)
}

func (fi FactoryImpl[databaseID]) ToDomain(propertyDatabaseModel Model[databaseID]) (*Property, error) {
	propertyDomainModel, err := fi.mapToDomain(fi.mapToDomainFunc, propertyDatabaseModel)
	if err != nil {
		return nil, err
	}
	return propertyDomainModel, fi.validate(propertyDomainModel)
}

func (fi FactoryImpl[databaseID]) ToDatabase(propertyDomainModel Property) (*Model[databaseID], error) {
	validationErr := fi.validate(&propertyDomainModel)
	if validationErr != nil {
		return nil, validationErr
	}
	propertyDatabaseModel, err := fi.mapToDatabase(fi.mapToDatabaseFunc, propertyDomainModel)
	if err != nil {
		return nil, err
	}
	return propertyDatabaseModel, nil
}
