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

// GetPropertyQuery : This is used to retrieve a property model.
type GetPropertyQuery struct {
	ID string `validate:"required"`
}

// GetPropertyHandler is a CQRS endpoint that handles a query to retrieve a property's information.
// It implements the QueryHandler interface for the GetPropertyQuery.
// The handler retrieves the property model from the database and returns it to the caller.
type GetPropertyHandler decorator.QueryHandler[GetPropertyQuery, *property.Property]

type GetPropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
}

// NewGetPropertyHandler creates a new instance of GetPropertyHandler,
// applying decorators for logging and validation.
func NewGetPropertyHandler(
	propRepo property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) GetPropertyHandler {
	if propRepo == nil {
		panic("nil property repository")
	}
	return decorator.ApplyQueryDecorators(
		GetPropertyHandlerImpl{
			repository: propRepo,
			validator:  validator,
		},
		logger,
		validator,
	)
}

// Handler method takes a context and returns a property model and an error.
func (guh GetPropertyHandlerImpl) Handle(c context.Context, cmd GetPropertyQuery,
) (*property.Property, error) {
	property, err := guh.repository.Get(c, cmd.ID)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	return property, nil
}
