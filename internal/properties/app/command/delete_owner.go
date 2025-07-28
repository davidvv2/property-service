package command

import (
	"context"

	"property-service/internal/properties/domain/owner"
	"property-service/pkg/decorator"
	"property-service/pkg/errors"
	"property-service/pkg/errors/codes"
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

// DeleteOwnerCommand : This is the delete owner request in a struct format.
type DeleteOwnerCommand struct {
	OwnerID string `validate:"required"`
}

// DeleteOwnerHandler is a CQRS endpoint that handles a command to delete an owner.
// It implements the CommandHandler interface for the DeleteOwnerCommand.
// The handler creates a delete an owner from the database.
type DeleteOwnerHandler decorator.CommandHandler[DeleteOwnerCommand]

type DeleteOwnerHandlerImpl struct {
	repository owner.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewDeleteOwnerHandler creates a new instance of DeleteOwnerHandler,
// applying necessary decorators for logging and validation.
func NewDeleteOwnerHandler(
	repository owner.Repository,
	logger log.Logger,
	validator *validator.Validate,
) DeleteOwnerHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		DeleteOwnerHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}

// Validate the delete owner command.
func (cph DeleteOwnerHandlerImpl) Handle(
	c context.Context, cmd DeleteOwnerCommand,
) error {
	if registerErr := cph.repository.Delete(
		c,
		cmd.OwnerID,
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}
