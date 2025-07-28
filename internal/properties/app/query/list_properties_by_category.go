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

// ListPropertiesByCategoryQuery : This is used to update the property profile.
type ListPropertiesByCategoryQuery struct {
	Category        string `validate:"required"`
	Sort            uint8  `validate:"required"`
	Search          uint8  `validate:"omitempty"`
	Limit           uint16 `validate:"required"`
	PaginationToken string `validate:"omitempty"`
}

// ListPropertiesByCategoryHandler is a CQRS endpoint that handles a command to retrieve a list of properties by category.
// It implements the QueryHandler interface for the ListPropertiesByCategoryQuery.
// The handler retrieves the property model from the database and returns it to the caller.
type ListPropertiesByCategoryHandler decorator.QueryHandler[ListPropertiesByCategoryQuery, *ListPropertiesByCategoryResult]

type ListPropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
}

// NewListPropertiesByCategoryHandler creates a new instance of ListPropertiesByCategoryHandler,
// applying decorators for logging and validation.
func NewListPropertiesByCategoryHandler(
	propRepo property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) ListPropertiesByCategoryHandler {
	if propRepo == nil {
		panic("nil property repository")
	}
	return decorator.ApplyQueryDecorators(
		ListPropertyHandlerImpl{
			repository: propRepo,
			validator:  validator,
		},
		logger,
		validator,
	)
}

// Handler method takes a context and returns a ListPropertiesByCategoryResult
// and an error.
func (guh ListPropertyHandlerImpl) Handle(c context.Context, cmd ListPropertiesByCategoryQuery,
) (*ListPropertiesByCategoryResult, error) {
	property, err := guh.repository.ListByCategory(
		c,
		cmd.Category,
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
	return &ListPropertiesByCategoryResult{
		Properties: property,
	}, nil
}

type ListPropertiesByCategoryResult struct {
	Properties []property.Property `json:"properties"`
}
