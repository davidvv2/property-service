package query

import (
	"context"

	"property-service/internal/properties/domain/property"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// ListPropertiesByOwnerQuery : This is used to update the property profile.
type ListPropertiesByOwnerQuery struct {
	Owner           string `validate:"required"`
	Sort            uint8  `validate:"required"`
	Search          uint8  `validate:"omitempty"`
	Limit           uint16 `validate:"required"`
	PaginationToken string `validate:"omitempty"`
	Server          string `validate:"required"`
}

// ListPropertiesByOwnerHandler is a CQRS endpoint that handles a command to retrieve a list of properties by category.
// It implements the QueryHandler interface for the ListPropertiesByOwnerQuery.
// The handler retrieves the property model from the database and returns it to the caller.
type ListPropertiesByOwnerHandler decorator.QueryHandler[ListPropertiesByOwnerQuery, *ListPropertiesByOwnerResult]

type ListPropertyByOwnerHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
}

// NewListPropertiesByOwnerHandler creates a new instance of ListPropertiesByOwnerHandler,
// applying decorators for logging and validation.
func NewListPropertiesByOwnerHandler(
	propRepo property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) ListPropertiesByOwnerHandler {
	if propRepo == nil {
		panic("nil property repository")
	}
	return decorator.ApplyQueryDecorators(
		ListPropertyByOwnerHandlerImpl{
			repository: propRepo,
			validator:  validator,
		},
		logger,
		validator,
	)
}

// Handler method takes a context and returns a ListPropertiesByOwnerResult
// and an error.
func (guh ListPropertyByOwnerHandlerImpl) Handle(c context.Context, cmd ListPropertiesByOwnerQuery,
) (*ListPropertiesByOwnerResult, error) {
	property, err := guh.repository.ListByOwner(
		c,
		cmd.Owner,
		cmd.Sort,
		cmd.Limit,
		cmd.PaginationToken,
		cmd.Search,
	)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	return &ListPropertiesByOwnerResult{
		Properties: property,
	}, nil
}

type ListPropertiesByOwnerResult struct {
	Properties []property.Property `json:"properties"`
}
