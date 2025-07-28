package property

import (
	"context"
)

// Repository :  handles all the database actions for the user profile.
type Repository interface {
	// New : property params.
	New(c context.Context, parms NewPropertyParams) (*Property, error)
	// Delete : Deletes a property by their id.
	Delete(c context.Context, ID string) error
	// Get : returns a single property by their id.
	Get(c context.Context, ID string) (*Property, error)
	// Update: updates a property.
	Update(c context.Context, id string, params UpdatePropertyParams) error

	ListByCategory(
		c context.Context,
		category string,
		sort uint8,
		limit uint16,
		paginationToken string,
		search uint8,
	) ([]Property, error)

	ListByOwner(
		c context.Context,
		ownerID string,
		sort uint8,
		limit uint16,
		paginationToken string,
		search uint8,
	) ([]Property, error)
}
