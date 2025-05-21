package property

import (
	"context"
)

// Repository :  handles all the database actions for the user profile.
type Repository interface {
	// New : property params.
	New(c context.Context, server string, parms NewPropertyParams) (*Property, error)
	// Delete : Deletes a property by their id.
	Delete(c context.Context, server string, ID string) error
	// Get : returns a single property by their id.
	Get(c context.Context, server string, ID string) (*Property, error)
	// Update: updates a property.
	Update(c context.Context, server string, id string, params UpdatePropertyParams) error

	ListByCategory(
		c context.Context,
		server string,
		category string,
		sort uint8,
		limit uint16,
		paginationToken string,
		search uint8,
	) ([]Property, error)
}
