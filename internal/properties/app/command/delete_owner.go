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

// DeleteOwnerCommand : This is the register request in a struct format.
type DeleteOwnerCommand struct {
	OwnerID string `validate:"required"`
	Server  string `validate:"required"`
}

// Validate the register User command.
func (cph DeleteOwnerHandlerImpl) Handle(
	c context.Context, cmd DeleteOwnerCommand,
) error {
	if registerErr := cph.repository.Delete(
		c,
		cmd.Server,
		cmd.OwnerID,
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}

// DeleteOwnerHandler is a CQRS endpoint that handles a command to retrieve a user's login attempt history.
// It implements the CommandHandler interface for the VerifyDeviceCommand.
// The handler retrieves the user's login attempt history from the database and returns it to the caller.
type DeleteOwnerHandler decorator.CommandHandler[DeleteOwnerCommand]

type DeleteOwnerHandlerImpl struct {
	repository owner.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewDeleteOwnerHandler : handles the login attempt query.
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
