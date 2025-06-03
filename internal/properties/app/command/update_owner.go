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

// UpdateOwnerCommand : This is the update owner request in a struct format.
type UpdateOwnerCommand struct {
	OwnerID   string `validate:"required"`
	Name      string
	Email     string
	Telephone string
	Server    string `validate:"required"`
}

// UpdateOwnerHandler is a CQRS endpoint that handles a command to update an owner's information.
// It implements the CommandHandler interface for the UpdateOwnerCommand.
// The handler updates an owner's information in the database.
type UpdateOwnerHandler decorator.CommandHandler[UpdateOwnerCommand]

type UpdateOwnerHandlerImpl struct {
	repository owner.Repository
	validator  *validator.Validate
	log        log.Logger
}

// NewUpdateOwnerHandler creates a new instance of UpdateOwnerHandler,
// applying necessary decorators for logging and validation.
func NewUpdateOwnerHandler(
	repository owner.Repository,
	logger log.Logger,
	validator *validator.Validate,
) UpdateOwnerHandler {
	if repository == nil {
		logger.Panic("nil repository")
	}
	return decorator.ApplyCommandDecorators(
		UpdateOwnerHandlerImpl{
			repository: repository,
			validator:  validator,
			log:        logger,
		},
		logger,
		validator,
	)
}

// Handle the update owner command.
func (cph UpdateOwnerHandlerImpl) Handle(
	c context.Context, cmd UpdateOwnerCommand,
) error {
	if registerErr := cph.repository.Update(
		c,
		cmd.Server,
		cmd.OwnerID,
		owner.UpdateOwnerParams{
			Telephone: cmd.Telephone,
			Email:     cmd.Email,
			Name:      cmd.Name,
		},
	); registerErr != nil {
		return errors.NewHandlerError(
			registerErr,
			codes.Internal,
		)
	}
	return nil
}
