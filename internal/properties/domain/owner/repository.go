package owner

import (
	"context"
)

// Repository :  handles all the database actions for the user profile.
type Repository interface {
	// New : property params.
	New(c context.Context, parms NewOwnerParams) (*Owner, error)
	// Delete : Deletes a property by their id.
	Delete(c context.Context, ID string) error
	// Get : returns a single property by their id.
	Get(c context.Context, ID string) (*Owner, error)
	// Update: updates a property.
	Update(
		c context.Context,
		ID string,
		params UpdateOwnerParams,
	) error
}
