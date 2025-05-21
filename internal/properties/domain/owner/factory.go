package owner

import (
	"property-service/pkg/errors"
	"property-service/pkg/helper/factory"
)

const (
	// Factory Config Constants.
	MaxSchemaVersion = 9999
)

type Factory[DatabaseID any] interface {
	New(
		property NewOwnerParams,
	) (*Owner, error)
	validate(p *Owner) error
	factory.Factory[Owner, Model[DatabaseID]]
}

// FactoryConfig is a struct for configuring the factory.
type FactoryConfig struct {
	SchemaVersion int
	/////TODO : add factory configs
}

// Validate validates the factory configuration and returns an error if it is invalid.
func (p FactoryConfig) Validate() error {
	switch {
	case p.SchemaVersion > MaxSchemaVersion:
		return errors.ErrMaxSchemaVersion
	case p.SchemaVersion <= 0:
		return errors.ErrMinSchemaVersion
	}
	return nil
}

type UpdateOwnerParams struct {
	Name      string
	Email     string
	Telephone string
}
