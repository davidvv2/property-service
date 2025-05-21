package adapters

import (
	"context"
	"time"

	"property-service/internal/properties/domain/owner"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Verify that OwnerRepositoryMongoImpl implements owner.Repository.
var _ owner.Repository = (*OwnerRepositoryMongoImpl)(nil)

type OwnerRepositoryMongoImpl struct {
	log   log.Logger
	owner database.FinderInserterUpdaterRemover[
		bson.M,
		bson.M,
		owner.Owner,
	]

	queryHelper database.QueryHelper[
		database.Query[string],
		database.TimeQuery[string],
		options.FindOptions,
		bson.M,
	]
	factory owner.Factory[primitive.ObjectID]
}

func NewMongoOwnerRepository(
	log log.Logger,
	owner database.FinderInserterUpdaterRemover[bson.M, bson.M, owner.Owner],
	factory owner.Factory[primitive.ObjectID],
) *OwnerRepositoryMongoImpl {
	return &OwnerRepositoryMongoImpl{
		log:         log,
		owner:       owner,
		queryHelper: database.NewMongoQueryHelper(),
		factory:     factory,
	}
}

func (p *OwnerRepositoryMongoImpl) New(
	ctx context.Context,
	server string,
	ownerParams owner.NewOwnerParams,
) (*owner.Owner, error) {
	p.log.Debug("Creating new owner")

	// Create a new owner using the factory
	newOwner, err := p.factory.New(ownerParams)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}

	// Insert the new owner into the database
	if _, err := p.owner.InsertOne(ctx, server, *newOwner); err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}

	return newOwner, nil
}

// Delete implements owner.Repository.
func (p *OwnerRepositoryMongoImpl) Delete(c context.Context, server string, ID string) error {
	p.log.Debug("Deleting owner with ID: %s", ID)
	count, err := p.owner.DeleteOneByID(c, server, ID)
	if err != nil {
		return errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	if count == 0 {
		return errors.NewHandlerError(
			err,
			codes.NotFound,
		)
	}
	return nil
}

// Get implements owner.Repository.
func (p *OwnerRepositoryMongoImpl) Get(c context.Context, server string, ID string) (*owner.Owner, error) {
	// TODO: Implement fetching a owner by ID from MongoDB.
	p.log.Debug("Fetching owner with ID: %s", ID)
	prop, getErr := p.owner.FindByID(c, server, ID)
	if getErr != nil {
		return nil, errors.NewHandlerError(
			getErr,
			codes.Internal,
		)
	}
	return prop, nil
}

// Update implements owner.Repository.
func (p *OwnerRepositoryMongoImpl) Update(c context.Context, server string, id string, params owner.UpdateOwnerParams) error {
	p.log.Debug("Updating owner with ID: %s", id)

	updateData := bson.M{}

	if params.Name != "" { // if pointer type, or check for a valid condition
		updateData["Name"] = params.Name
	}
	if params.Telephone != "" {
		updateData["Telephone"] = params.Telephone
	}
	if params.Email != "" {
		updateData["Email"] = params.Email
	}
	if len(updateData) == 0 {
		return nil // nothing to update
	}

	updateData["Metadata.UpdatedAt"] = primitive.NewDateTimeFromTime(time.Now())
	updateFields := bson.M{
		"$set": updateData,
	}

	err := p.owner.UpdateOneByID(c, server, id, updateFields)
	if err != nil {
		return errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	return nil

}
