package adapters

import (
	"context"

	"property-service/internal/properties/domain/property"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/database"
	"property-service/pkg/infrastructure/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Verify that PropertyRepositoryMongoImpl implements property.Repository.
var _ property.Repository = (*PropertyRepositoryMongoImpl)(nil)

type PropertyRepositoryMongoImpl struct {
	log      log.Logger
	property database.FinderInserterUpdaterRemover[
		bson.M,
		bson.M,
		property.Property,
	]

	queryHelper database.QueryHelper[
		database.Query[string],
		database.TimeQuery[string],
		options.FindOptions,
		bson.M,
	]
	paginationHelper database.PaginationHelper
	factory          property.Factory[primitive.ObjectID]
	aggregator       database.Grouper[mongo.Pipeline, property.Property]
}

func NewMongoPropertyRepository(
	log log.Logger,
	property database.FinderInserterUpdaterRemover[bson.M, bson.M, property.Property],
	factory property.Factory[primitive.ObjectID],
	aggregator database.Grouper[mongo.Pipeline, property.Property],
) *PropertyRepositoryMongoImpl {
	return &PropertyRepositoryMongoImpl{
		log:              log,
		property:         property,
		queryHelper:      database.NewMongoQueryHelper(),
		paginationHelper: &database.PaginationHelperMongoImpl{},
		factory:          factory,
		aggregator:       aggregator,
	}
}

func (p *PropertyRepositoryMongoImpl) New(
	ctx context.Context,
	server string,
	propertyParams property.NewPropertyParams,
) (*property.Property, error) {
	p.log.Debug("Creating new property")

	// Create a new property using the factory
	newProperty, err := p.factory.New(propertyParams)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}

	// Insert the new property into the database
	if _, err := p.property.InsertOne(ctx, server, *newProperty); err != nil {
		return nil, errors.NewHandlerError(err,
			codes.Internal,
		)
	}

	return newProperty, nil
}

// Delete implements property.Repository.
func (p *PropertyRepositoryMongoImpl) Delete(c context.Context, server string, ID string) error {
	p.log.Debug("Deleting property with ID: %s", ID)
	count, err := p.property.DeleteOneByID(c, server, ID)
	if err != nil {
		return errors.NewHandlerError(err,
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

// Get implements property.Repository.
func (p *PropertyRepositoryMongoImpl) Get(c context.Context, server string, ID string) (*property.Property, error) {
	// TODO: Implement fetching a property by ID from MongoDB.
	p.log.Debug("Fetching property with ID: %s", ID)
	prop, getErr := p.property.FindByID(c, server, ID)
	if getErr != nil {
		return nil, errors.NewHandlerError(getErr,
			codes.Internal,
		)
	}
	return prop, nil
}

// Update implements property.Repository.
func (p *PropertyRepositoryMongoImpl) Update(
	c context.Context,
	server string,
	id string,
	params property.UpdatePropertyParams,
) error {
	p.log.Debug("Updating property with ID: %s", id)

	updateData := bson.M{}

	if params.Available != nil { // if pointer type, or check for a valid condition
		updateData["Available"] = *params.Available
	}
	if !params.AvailableDate.IsZero() {
		updateData["AvailableDate"] = params.AvailableDate
	}
	if params.Description != "" {
		updateData["Description"] = params.Description
	}
	if params.Title != "" {
		updateData["Title"] = params.Title
	}
	if params.Category != "" {
		updateData["Category"] = params.Category
	}
	if params.Address != "" {
		updateData["Address"] = params.Address
	}
	if params.SaleType != 0 {
		updateData["SaleType"] = params.SaleType
	}

	if len(updateData) == 0 {
		return nil // nothing to update
	}
	updateFields := bson.M{
		"$set": updateData,
	}
	err := p.property.UpdateOneByID(c, server, id, updateFields)
	if err != nil {
		return errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	return nil

}

// ListByCategory implements property.Repository.
func (p *PropertyRepositoryMongoImpl) ListByCategory(
	c context.Context,
	server string,
	category string,
	sort uint8,
	limit uint16,
	paginationToken string,
	search uint8,
) ([]property.Property, error) {
	sortSpec := bson.D{
		{Key: "Title", Value: sort},
	}
	filter, err := p.paginationHelper.TextPaginationHelper(
		"default",
		"Category",
		category,
		sortSpec,
		search,
		paginationToken,
	)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}

	res, aggErr := p.aggregator.Aggregate(
		c,
		server,
		mongo.Pipeline{
			filter,
			bson.D{{Key: "$limit", Value: limit}},
			bson.D{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "OwnerID", Value: 1},
				{Key: "Description", Value: 1},
				{Key: "Title", Value: 1},
				{Key: "Category", Value: 1},
				{Key: "Available", Value: 1},
				{Key: "AvailableDate", Value: 1},
				{Key: "Address", Value: 1},
				{Key: "SaleType", Value: 1},
				{Key: "PaginationToken", Value: bson.D{{Key: "$meta", Value: "searchSequenceToken"}}},
			}}}},
	)
	if aggErr != nil {
		p.log.Debug("Error in aggregation: %v", aggErr)
		return nil, errors.NewHandlerError(
			aggErr,
			codes.Internal,
		)
	}

	finalRes, getErr := res.GetAll(c)
	p.log.InfoWithFields("Fetched properties",
		log.Fields{
			"count":           len(*finalRes),
			"server":          server,
			"category":        category,
			"sort":            sort,
			"limit":           limit,
			"paginationToken": paginationToken,
			"properties":      *finalRes,
		},
	)
	if getErr != nil {
		p.log.Debug("Error in getting all results: %v", getErr)
		return nil, errors.NewHandlerError(
			getErr,
			codes.Internal,
		)
	}
	return *finalRes, nil
}
