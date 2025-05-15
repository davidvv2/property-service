package owner

import "time"

type SaleType uint8

type Model[ID any] struct {
	ID        ID            `bson:"_id" validate:"required"`
	Name      string        `bson:"Name" validate:"required"`
	Email     string        `bson:"Email" validate:"required,email"`
	Telephone string        `bson:"Telephone" validate:"required"`
	Metadata  MetadataModel `bson:"Metadata" validate:"required"`
}

type MetadataModel struct {
	CreatedAt time.Time `bson:"CreatedAt"`
	UpdatedAt time.Time `bson:"UpdatedAt"`
}

func MapModelToOwner[Old any](
	mappingFunc func(Old) (string, error),
	oldOwner Model[Old],
) (*Owner, error) {
	// Map IDs
	ownerID, ownerIDErr := mappingFunc(oldOwner.ID)
	if ownerIDErr != nil {
		return nil, ownerIDErr
	}
	return &Owner{
		id:        ownerID,
		name:      oldOwner.Name,
		email:     oldOwner.Email,
		telephone: oldOwner.Telephone,
		metadata: Metadata{
			createdAt: oldOwner.Metadata.CreatedAt,
			updatedAt: oldOwner.Metadata.UpdatedAt,
		},
	}, ownerIDErr
}

// Owner : This domain model contains a property voucher model.
type Owner struct {
	id        string   `validate:"required"`
	name      string   `validate:"required"`
	email     string   `validate:"required"`
	telephone string   `validate:"required"`
	metadata  Metadata `validate:"required"`
}

type Metadata struct {
	createdAt time.Time `bson:"CreatedAt"`
	updatedAt time.Time `bson:"UpdatedAt"`
}

func MapOwnerToModel[New any](
	mappingFunc func(string) (New, error),
	oldOwner Owner,
) (*Model[New], error) {
	// Map IDs
	ownerID, ownerIDErr := mappingFunc(oldOwner.id)
	if ownerIDErr != nil {
		return nil, ownerIDErr
	}

	return &Model[New]{
		ID:        ownerID,
		Name:      oldOwner.name,
		Email:     oldOwner.email,
		Telephone: oldOwner.telephone,
		Metadata: MetadataModel{
			CreatedAt: oldOwner.metadata.createdAt,
			UpdatedAt: oldOwner.metadata.updatedAt,
		},
	}, nil
}
