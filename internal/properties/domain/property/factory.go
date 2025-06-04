package property

import (
	"property-service/pkg/address"
	"property-service/pkg/helper/factory"
	"time"
)

const (
	// Factory Config Constants.
	MaxSchemaVersion = 9999
)

type Factory[DatabaseID any] interface {
	New(
		property NewPropertyParams,
	) (*Property, error)
	validate(p *Property) error
	factory.Factory[Property, Model[DatabaseID]]
}

// FactoryConfig is a struct for configuring the factory.
type FactoryConfig struct {
	SchemaVersion int
	/////TODO : add factory configs
}

// Validate validates the factory configuration and returns an error if it is invalid.
func (p FactoryConfig) Validate() error {
	return nil
}

type UpdatePropertyParams struct {
	AvailableDate time.Time
	Description   string
	Title         string
	Category      string
	Address       address.Address
	SaleType      uint8
}
