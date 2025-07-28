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

// DeletePropertyCommand : This is the delete property request in a struct format.
type DeletePropertyCommand struct {
	PropertyID string `validate:"required"`
}

// DeletePropertyHandler is a CQRS endpoint that handles a command to delete a property.
// It implements the CommandHandler interface for the DeletePropertyCommand.
// This handler is used to delete a property from the database.
type DeletePropertyHandler decorator.CommandHandler[DeletePropertyCommand]

type DeletePropertyHandlerImpl struct {
	repository property.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewDeletePropertyHandler creates a new instance of DeletePropertyHandler,
// applying necessary decorators for logging and validation.
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

// Validate the delete property command.
func (cph DeletePropertyHandlerImpl) Handle(
	c context.Context, cmd DeletePropertyCommand,
) error {
	if registerErr := cph.repository.Delete(
		c,
		cmd.PropertyID,
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}
