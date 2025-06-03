package query

import (
	"context"

	"property-service/internal/properties/domain/owner"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// GetOwnerQuery : This is used to retrieve the owner profile.
type GetOwnerQuery struct {
	ID     string `validate:"required"`
	Server string `validate:"required"`
}

// GetOwnerHandler is a CQRS endpoint that handles a command to retrieve a owner's profile.
// The handler retrieves the owner model from the database and returns it to the caller.
type GetOwnerHandler decorator.QueryHandler[GetOwnerQuery, *owner.Owner]

type getOwnerHandlerImpl struct {
	repository owner.Repository
	validator  *validator.Validate
}

// NewGetOwnerHandler creates a new instance of GetOwnerHandler,
// applying decorators for logging and validation.
func NewGetOwnerHandler(
	propRepo owner.Repository,
	logger log.Logger,
	validator *validator.Validate,
) GetOwnerHandler {
	if propRepo == nil {
		panic("nil owner repository")
	}
	return decorator.ApplyQueryDecorators(
		getOwnerHandlerImpl{
			repository: propRepo,
			validator:  validator,
		},
		logger,
		validator,
	)
}

// Handler method takes a context and returns an owner model and an error.
func (guh getOwnerHandlerImpl) Handle(c context.Context, cmd GetOwnerQuery,
) (*owner.Owner, error) {
	owner, err := guh.repository.Get(c, cmd.Server, cmd.ID)
	if err != nil {
		return nil, errors.NewHandlerError(
			err,
			codes.Aborted,
		)
	}
	return owner, nil
}
