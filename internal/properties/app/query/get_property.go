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

// GetPropertyQuery : This is used to update the property profile.
type GetPropertyQuery struct {
	ID     string `validate:"required"`
	Server string `validate:"required"`
}

// GetPropertyHandler is a CQRS endpoint that handles a command to retrieve a property's login attempt history.
// It implements the QueryHandler interface for the LoginAttemptQuery.
// The handler retrieves the property model from the database and returns it to the caller.
type GetPropertyHandler decorator.QueryHandler[GetPropertyQuery, *property.Property]

type GetPropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
}

// NewGetPropertyHandler : handles the get property attempt query.
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

// Handler method takes a context and returns a flat buffer response
// and an error.
func (guh GetPropertyHandlerImpl) Handle(c context.Context, cmd GetPropertyQuery,
) (*property.Property, error) {
	property, err := guh.repository.Get(c, cmd.Server, cmd.ID)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Internal,
		)
	}
	return property, nil
}
