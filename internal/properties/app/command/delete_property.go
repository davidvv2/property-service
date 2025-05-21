package command

import (
	"context"

	"property-service/internal/properties/domain/property"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// DeletePropertyCommand : This is the register request in a struct format.
type DeletePropertyCommand struct {
	PropertyID string `validate:"required"`
	Server     string `validate:"required"`
}

// Validate the register User command.
func (cph DeletePropertyHandlerImpl) Handle(
	c context.Context, cmd DeletePropertyCommand,
) error {
	if registerErr := cph.repository.Delete(
		c,
		cmd.Server,
		cmd.PropertyID,
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}

// DeletePropertyHandler is a CQRS endpoint that handles a command to retrieve a user's login attempt history.
// It implements the CommandHandler interface for the VerifyDeviceCommand.
// The handler retrieves the user's login attempt history from the database and returns it to the caller.
type DeletePropertyHandler decorator.CommandHandler[DeletePropertyCommand]

type DeletePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewDeletePropertyHandler : handles the login attempt query.
func NewDeletePropertyHandler(
	repository property.Repository,
	logger log.Logger,
	validator *validator.Validate,
) DeletePropertyHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		DeletePropertyHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}
