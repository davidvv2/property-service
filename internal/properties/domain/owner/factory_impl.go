package owner

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
	mapToDomain       func(mapper func(databaseID) (string, error), his Model[databaseID]) (*Owner, error)
	mapToDatabase     func(mapper func(string) (databaseID, error), his Owner) (*Model[databaseID], error)
	v                 *validator.Validate // validator used for validating the factory configuration
	fc                FactoryConfig       // configuration for the factory
}

// NewFactory creates a new Factory with the given configuration and returns an error if the configuration is invalid.
func NewFactory[databaseID comparable](
	fc FactoryConfig, v *validator.Validate, newID func() string,
	mappingFunc func(databaseID) (string, error),
	mapHistory func(mappingFunc func(databaseID) (string, error), databaseModel Model[databaseID]) (*Owner, error),
	mapToDatabaseFunc func(string) (databaseID, error),
	mapToDatabase func(mappingFunc func(string) (databaseID, error), domainModel Owner) (*Model[databaseID], error),
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
	mapHistory func(mappingFunc func(databaseID) (string, error), databaseModel Model[databaseID]) (*Owner, error),
	mapToDatabaseFunc func(string) (databaseID, error),
	mapToDatabase func(mappingFunc func(string) (databaseID, error), domainModel Owner) (*Model[databaseID], error),
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

func (fi FactoryImpl[databaseID]) validate(la *Owner) error {
	return fi.v.Struct(la)
}

type NewOwnerParams struct {
	ID        string `validate:"required"`
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Telephone string `validate:"required"`
}

func (fi FactoryImpl[databaseID]) New(
	owner NewOwnerParams,
) (*Owner, error) {
	ownerModel := &Owner{
		id:        owner.ID,
		name:      owner.Name,
		email:     owner.Email,
		telephone: owner.Telephone,
		metadata: Metadata{
			createdAt: time.Now(),
			updatedAt: time.Time{},
		},
	}
	return ownerModel, fi.validate(ownerModel)
}

func (fi FactoryImpl[databaseID]) ToDomain(ownerDatabaseModel Model[databaseID]) (*Owner, error) {
	ownerDomainModel, err := fi.mapToDomain(fi.mapToDomainFunc, ownerDatabaseModel)
	if err != nil {
		return nil, err
	}
	return ownerDomainModel, fi.validate(ownerDomainModel)
}

func (fi FactoryImpl[databaseID]) ToDatabase(ownerDomainModel Owner) (*Model[databaseID], error) {
	validationErr := fi.validate(&ownerDomainModel)
	if validationErr != nil {
		return nil, validationErr
	}
	ownerDatabaseModel, err := fi.mapToDatabase(fi.mapToDatabaseFunc, ownerDomainModel)
	if err != nil {
		return nil, err
	}
	return ownerDatabaseModel, nil
}
