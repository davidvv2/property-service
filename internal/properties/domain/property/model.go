package property

import "time"

type SaleType uint8

const (
	Unknown SaleType = iota // 0: unknown
	ForRent                 // 1: for rent
	ForSale                 // 2: for sale
	ForBoth                 // 3: both
)

type Model[ID any] struct {
	ID            ID            `bson:"_id" validate:"required,len=24,hexadecimal"`
	OwnerID       ID            `bson:"OwnerID" validate:"required,len=24,hexadecimal"`
	Category      string        `bson:"Category" validate:"required"`
	Description   string        `bson:"Description" validate:"required"`
	Title         string        `bson:"Title" validate:"required"`
	Metadata      MetadataModel `bson:"Metadata" validate:"required"`
	Available     bool          `bson:"Available" validate:"required"`
	AvailableDate time.Time     `bson:"AvailableDate" validate:"required"`
	Address       string        `bson:"Address" validate:"required"`
	SaleType      SaleType      `bson:"SaleType" validate:"gte=0,lte=3"`
}

type MetadataModel struct {
	CreatedAt time.Time `bson:"CreatedAt"`
	UpdatedAt time.Time `bson:"UpdatedAt"`
}

func MapModelToProperty[Old any](
	mappingFunc func(Old) (string, error),
	oldProperty Model[Old],
) (*Property, error) {
	// Map IDs
	propertyID, err := mappingFunc(oldProperty.ID)
	if err != nil {
		return nil, err
	}
	ownerID, ownerIDErr := mappingFunc(oldProperty.OwnerID)
	if ownerIDErr != nil {
		return nil, ownerIDErr
	}
	return &Property{
		id:          propertyID,
		ownerID:     ownerID,
		category:    oldProperty.Category,
		description: oldProperty.Description,
		title:       oldProperty.Title,
		metadata: Metadata{
			createdAt: oldProperty.Metadata.CreatedAt,
			updatedAt: oldProperty.Metadata.UpdatedAt,
		},
		available:     oldProperty.Available,
		availableDate: oldProperty.AvailableDate,
		address:       oldProperty.Address,
		saleType:      uint8(oldProperty.SaleType),
	}, err
}

// Property : This domain model contains a property voucher model.
type Property struct {
	id            string    `validate:"required"`
	ownerID       string    `validate:"required"`
	category      string    `validate:"required"`
	description   string    `validate:"required"`
	title         string    `validate:"required"`
	metadata      Metadata  `validate:"required"`
	available     bool      `validate:"required"`
	availableDate time.Time `validate:"required"`
	address       string    `validate:"required"`
	saleType      uint8     `validate:"required"`
}

type Metadata struct {
	createdAt time.Time `bson:"CreatedAt"`
	updatedAt time.Time `bson:"UpdatedAt"`
}

func MapPropertyToModel[New any](
	mappingFunc func(string) (New, error),
	oldProperty Property,
) (*Model[New], error) {
	// Map IDs
	propertyID, err := mappingFunc(oldProperty.id)
	if err != nil {
		return nil, err
	}
	ownerID, ownerIDErr := mappingFunc(oldProperty.id)
	if ownerIDErr != nil {
		return nil, ownerIDErr
	}

	return &Model[New]{
		ID:      propertyID,
		OwnerID: ownerID,

		Category:    oldProperty.category,
		Description: oldProperty.description,
		Title:       oldProperty.title,
		Metadata: MetadataModel{
			CreatedAt: oldProperty.metadata.createdAt,
			UpdatedAt: oldProperty.metadata.updatedAt,
		},
		Available:     oldProperty.available,
		AvailableDate: oldProperty.availableDate,
		Address:       oldProperty.address,
		SaleType:      SaleType(oldProperty.saleType),
	}, err
}
