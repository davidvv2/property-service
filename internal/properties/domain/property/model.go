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
	ID              ID            `bson:"_id" validate:"required,len=24,hexadecimal"`
	OwnerID         ID            `bson:"OwnerID" validate:"required,len=24,hexadecimal"`
	Category        string        `bson:"Category" validate:"required"`
	Description     string        `bson:"Description" validate:"required"`
	Title           string        `bson:"Title" validate:"required"`
	Metadata        MetadataModel `bson:"Metadata" validate:"required"`
	Available       bool          `bson:"Available" validate:"required"`
	AvailableDate   time.Time     `bson:"AvailableDate" validate:"required"`
	Address         string        `bson:"Address" validate:"required"`
	SaleType        SaleType      `bson:"SaleType" validate:"gte=0,lte=3"`
	PaginationToken string        `bson:"PaginationToken,omitempty" validate:"omitempty"`
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
		ID:          propertyID,
		OwnerID:     ownerID,
		Category:    oldProperty.Category,
		Description: oldProperty.Description,
		Title:       oldProperty.Title,
		Metadata: Metadata{
			createdAt: oldProperty.Metadata.CreatedAt,
			updatedAt: oldProperty.Metadata.UpdatedAt,
		},
		Available:       oldProperty.Available,
		AvailableDate:   oldProperty.AvailableDate,
		Address:         oldProperty.Address,
		SaleType:        uint8(oldProperty.SaleType),
		PaginationToken: oldProperty.PaginationToken,
	}, err
}

// Property : This domain model contains a property voucher model.
type Property struct {
	ID              string    `json:"id" validate:"required"`
	OwnerID         string    `json:"ownerID" validate:"required"`
	Category        string    `json:"category" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Title           string    `json:"title" validate:"required"`
	Metadata        Metadata  `json:"metadata" validate:"required"`
	Available       bool      `json:"available" validate:"required"`
	AvailableDate   time.Time `json:"availableDate" validate:"required"`
	Address         string    `json:"address" validate:"required"`
	SaleType        uint8     `json:"saleType" validate:"required"`
	PaginationToken string    `json:"paginationToken,omitempty" validate:"omitempty"`
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
	propertyID, err := mappingFunc(oldProperty.ID)
	if err != nil {
		return nil, err
	}
	ownerID, ownerIDErr := mappingFunc(oldProperty.OwnerID)
	if ownerIDErr != nil {
		return nil, ownerIDErr
	}

	return &Model[New]{
		ID:      propertyID,
		OwnerID: ownerID,

		Category:    oldProperty.Category,
		Description: oldProperty.Description,
		Title:       oldProperty.Title,
		Metadata: MetadataModel{
			CreatedAt: oldProperty.Metadata.createdAt,
			UpdatedAt: oldProperty.Metadata.updatedAt,
		},
		Available:       oldProperty.Available,
		AvailableDate:   oldProperty.AvailableDate,
		Address:         oldProperty.Address,
		SaleType:        SaleType(oldProperty.SaleType),
		PaginationToken: oldProperty.PaginationToken,
	}, err
}
